package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
	"ranksrv/proto/pb"
	"ranksrv/rank/persistence"
	"time"
)

var (
	TimeOutError = errors.New("request timeout")
	QueueIsFull  = errors.New("job queue is full")
)

type PartRank struct {
	running   bool // 运行标志位,多线程竞争不敏感
	rankInfo  *pb.RankInfo
	size      int
	cachePath string
	pstTick   *time.Ticker                //数据持久化定时器
	pstCache  *persistence.SequenceCache  // 持久化缓存
	rtCache   map[string]*pb.RankUnitData // 实时缓存
	rArray    *RegularArray               // 基础数组
	jobQueue  chan interface{}            //工作队列
	close     chan int
}

func NewPartRank(rankInfo *pb.RankInfo) *PartRank {
	return &PartRank{
		running:  false,
		rankInfo: rankInfo,
		size:     int(rankInfo.TopNum) * 120 / 100,
		jobQueue: make(chan interface{}, 512),
		close:    make(chan int),
	}
}

func (w *PartRank) Start(cacheDir string) (err error) {
	if w.running {
		return
	}
	if w.rankInfo.ScoreNum <= 0 {
		err = ScoreNumError
		return
	}
	w.rArray = NewRegularArray(w.size)
	w.rtCache = make(map[string]*pb.RankUnitData, w.size)
	// 读取持久化缓存取数据
	w.cachePath = cacheDir + w.rankInfo.Name
	w.pstCache = persistence.NewSequenceCache(w.cachePath)
	if err = w.pstCache.Init(); err != nil {
		return
	}
	data, err := w.pstCache.Read()
	if err != nil {
		return
	}
	for _, v := range data {
		unitData := &pb.RankUnitData{}
		if err = proto.Unmarshal(v,unitData); err != nil {
			return
		}
		w.rArray.Append(unitData)
		w.rtCache[unitData.UniqueID] = unitData
	}
	// 启动工作线程
	go w.doWork()
	// 启动数据持久化线程
	go w.doPersistence()
	w.running = true
	return
}

func (w *PartRank) RankInfo() *pb.RankInfo {
	return w.rankInfo
}

func (w *PartRank) doWork() {
	for job := range w.jobQueue {
		switch job.(type) {
		case *UpdateJob:
			w.doUpdate(job.(*UpdateJob))
		case *DeleteJob:
			w.doDelete(job.(*DeleteJob))
		case *QueryRankJob:
			w.doQueryRank(job.(*QueryRankJob))
		case *QueryTopJob:
			w.doQueryTop(job.(*QueryTopJob))
		case *QueryByScoreJob:
			w.doQueryByScore(job.(*QueryByScoreJob))
		default:
			fmt.Println("Don't define job")
		}
	}
}

func (w *PartRank) doPersistence() {
	w.pstTick = time.NewTicker(time.Second * time.Duration(w.rankInfo.FlushInterval))
	for {
		select {
		case <-w.pstTick.C:
			topRank, err := w.QueryTop(w.rankInfo.TopNum)
			if err != nil {
				continue
			}
			data := make([][]byte, 0, len(topRank))
			var buf []byte
			for _, v := range topRank {
				buf, err = proto.Marshal(v)
				if err != nil {
					fmt.Println("Do persistence marshal failed", err)
					break
				}
				data = append(data, buf)
			}
			if err != nil { //如果编码处理有错误，跳过写入
				continue
			}
			if err = w.pstCache.ReWrite(data); err != nil {
				fmt.Println("Write file failed", err)
			}
		case <-w.close:
			return
		}
	}
}

func (w *PartRank) putJob(job interface{}) (err error) {
	if !w.running {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("job queue is full")
		err = QueueIsFull
		return
	case w.jobQueue <- job:
		return
	}
	return
}

func (w *PartRank) Update(data *pb.RankUnitData) {
	if len(data.Score) != int(w.rankInfo.ScoreNum) {
		return
	}
	if data.Score[0] < w.rankInfo.ScoreMinLimit {
		return
	}
	updateJob := &UpdateJob{
		Data: data,
	}
	w.putJob(updateJob)
}

func (w *PartRank) doUpdate(job *UpdateJob) {
	newPos, ok := w.rArray.Find(job.Data)
	if ok {
		return
	} else {
		data, ok := w.rtCache[job.Data.UniqueID]
		if ok {
			oldPos, ok := w.rArray.Find(data)
			if ok {
				w.rArray.Update(oldPos, newPos, job.Data)
				w.rtCache[job.Data.UniqueID] = job.Data
			}
		} else {
			if newPos == w.rArray.Size() {
				if w.rArray.Append(job.Data) {
					w.rtCache[job.Data.UniqueID] = job.Data
				}
			} else {
				endData := w.rArray.Insert(newPos, job.Data)
				if endData != nil {
					delete(w.rtCache, endData.UniqueID)
				}
				w.rtCache[job.Data.UniqueID] = job.Data
			}
		}
	}
}

func (w *PartRank) Delete(uniqueID string) {
	deleteJob := &DeleteJob{
		UniqueID: uniqueID,
	}
	w.putJob(deleteJob)
	return
}

func (w *PartRank) doDelete(job *DeleteJob) {
	data, ok := w.rtCache[job.UniqueID]
	if ok {
		pos, ok := w.rArray.Find(data)
		if ok {
			w.rArray.Delete(pos)
			delete(w.rtCache, job.UniqueID)
		}
	}
	return
}

func (w *PartRank) QueryRank(uniqueID string) (rankInfo *pb.UnitRankInfo, err error) {
	rspChannel := make(chan *pb.UnitRankInfo, 1)
	job := &QueryRankJob{
		UniqueID:   uniqueID,
		RspChannel: rspChannel,
	}
	if err = w.putJob(job); err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	select {
	case <-ctx.Done():
		err = TimeOutError
		return
	case rankInfo = <-rspChannel:
		return
	}
	return
}

func (w *PartRank) doQueryRank(job *QueryRankJob) {
	data, ok := w.rtCache[job.UniqueID]
	var rankInfo *pb.UnitRankInfo
	if ok {
		pos, ok := w.rArray.Find(data)
		if ok {
			rankInfo = &pb.UnitRankInfo{
				Rank:  uint32(pos),
				Score: data.Score,
			}
		}
	}
	job.RspChannel <- rankInfo
}

func (w *PartRank) QueryTop(reqNum uint32) (topList []*pb.RankUnitData, err error) {
	if reqNum > w.rankInfo.TopNum {
		reqNum = w.rankInfo.TopNum
	}
	rspChannel := make(chan []*pb.RankUnitData, 1)
	job := &QueryTopJob{
		reqNum:     int(reqNum),
		RspChannel: rspChannel,
	}
	if err = w.putJob(job); err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	select {
	case <-ctx.Done():
		err = TimeOutError
		return
	case topList = <-rspChannel:
		return
	}
	return
}

func (w *PartRank) doQueryTop(job *QueryTopJob) {
	top := w.rArray.GetTop(job.reqNum)
	job.RspChannel <- top
}

func (w *PartRank) QueryByScore(score []uint64) (rank uint32, err error) {
	rspChannel := make(chan int, 1)
	job := &QueryByScoreJob{
		Score:      score,
		RspChannel: rspChannel,
	}
	if err = w.putJob(job); err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	select {
	case <-ctx.Done():
		err = TimeOutError
		return
	case pos := <-rspChannel:
		rank = uint32(pos)
		return
	}
	return
}

func (w *PartRank) doQueryByScore(job *QueryByScoreJob) (rank uint32) {
	pos, _ := w.rArray.FindByScore(job.Score)
	job.RspChannel <- pos
	return
}

func (w *PartRank) Stop() {
	if !w.running {
		return
	}
	w.running = false
	w.pstTick.Stop()
	w.close <- 1 //等待数据写入线程退出再close文件
	w.pstCache.Close()
	close(w.jobQueue)
}

func (w *PartRank) Close() {
	w.Stop()
	os.Remove(w.cachePath)
}
