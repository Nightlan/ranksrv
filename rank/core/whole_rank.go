package core

import (
	"errors"
	"fmt"
	"os"
	"ranksrv/proto/pb"
	"ranksrv/rank/persistence"
	"sort"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var (
	UniqueNotExist = errors.New("unique id not exist")
	ScoreNumError  = errors.New("score num error")
)

type WholeRank struct {
	running    bool // 运行标志位,多线程竞争不敏感
	rankInfo   *pb.RankInfo
	cachePath  string
	totalNum   uint32 // 排行榜总数
	updateTick *time.Ticker
	pstCache   *persistence.FormatCache // 持久化缓存
	rtCache    sync.Map                 // 实时缓存
	rankData   unsafe.Pointer           // 排名数据,类型:map[string]*pb.UnitRankInfo
	rank       unsafe.Pointer           // top榜单数据,类型:[]*pb.RankUnitData
	close      chan int
}

func NewWholeRank(rankInfo *pb.RankInfo) *WholeRank {
	return &WholeRank{
		running:  false,
		rankInfo: rankInfo,
		totalNum: 0,
		close:    make(chan int),
	}
}

func (w *WholeRank) Start(cacheDir string) (err error) {
	if w.running {
		return
	}
	if w.rankInfo.ScoreNum <= 0 {
		err = ScoreNumError
		return
	}
	// 读取持久化缓存取数据
	w.cachePath = cacheDir + w.rankInfo.Name
	w.pstCache = persistence.NewFormatCache(w.cachePath)
	if err = w.pstCache.Init(); err != nil {
		return
	}
	data, err := w.pstCache.Read()
	if err != nil {
		return
	}
	for _, v := range data {
		rankData := &CacheUnitData{
			RankUnitData: v,
			Update:       true,
		}
		w.rtCache.Store(rankData.UniqueID, rankData)
	}
	// 初次启动重新生成排行数据
	w.calculateRank()
	// 启动定时计算
	go w.doWork()
	w.running = true
	return
}

func (w *WholeRank) RankInfo() *pb.RankInfo {
	return w.rankInfo
}

func (w *WholeRank) doWork() {
	w.updateTick = time.NewTicker(time.Second * time.Duration(w.rankInfo.FlushInterval))
	for {
		select {
		case <-w.updateTick.C:
			if err := w.calculateRank(); err != nil {
				fmt.Println("Calculate rank failed", err)
			}
		case <-w.close:
			return
		}
	}
}

func (w *WholeRank) calculateRank() (err error) {
	// 数据排序
	totalNum := atomic.LoadUint32(&w.totalNum)
	rank := make([]*pb.RankUnitData, 0, totalNum)
	updateList := make([]*pb.RankUnitData, 0)
	w.rtCache.Range(func(key, value interface{}) bool {
		data := value.(*CacheUnitData)
		rank = append(rank, data.RankUnitData)
		if data.Update {
			updateList = append(updateList, data.RankUnitData)
			data.Update = false
		}
		return true
	})
	sort.Slice(rank, func(i, j int) bool {
		return UnitDataLess(rank[i], rank[j])
	})
	// 更新排名和榜单
	rankData := make(map[string]*pb.UnitRankInfo, totalNum)
	for i, v := range rank {
		data := &pb.UnitRankInfo{
			Rank:  uint32(i) + 1,
			Score: v.Score,
		}
		rankData[v.UniqueID] = data
	}
	atomic.StorePointer(&w.rank, unsafe.Pointer(&rank))
	atomic.StorePointer(&w.rankData, unsafe.Pointer(&rankData))

	// 数据持久化
	if err = w.pstCache.Write(updateList); err != nil {
		return
	}
	return
}

func (w *WholeRank) Update(data *pb.RankUnitData) {
	if len(data.Score) != int(w.rankInfo.ScoreNum) {
		return
	}
	if data.Score[0] < w.rankInfo.ScoreMinLimit {
		return
	}
	if _, ok := w.rtCache.Load(data.UniqueID); !ok {
		atomic.AddUint32(&w.totalNum, 1)
	}
	rankData := &CacheUnitData{
		RankUnitData: data,
		Update:       true,
	}
	w.rtCache.Store(rankData.UniqueID, rankData)
}

func (w *WholeRank) Delete(uniqueID string) {
	w.rtCache.Delete(uniqueID)
	return
}

func (w *WholeRank) QueryRank(uniqueID string) (rankInfo *pb.UnitRankInfo, err error) {
	pointer := atomic.LoadPointer(&w.rankData)
	rankData := (*map[string]*pb.UnitRankInfo)(pointer)
	info, ok := (*rankData)[uniqueID]
	if !ok {
		err = UniqueNotExist
	}
	rankInfo = info
	return
}

func (w *WholeRank) QueryTop(reqNum uint32) (topList []*pb.RankUnitData, err error) {
	pointer := atomic.LoadPointer(&w.rank)
	topRank := (*[]*pb.RankUnitData)(pointer)
	if int(reqNum) > len(*topRank) {
		reqNum = uint32(len(*topRank))
	}
	topList = (*topRank)[:reqNum]
	return
}

func (w *WholeRank) QueryByScore(score []uint64) (rank uint32, err error) {
	if len(score) != int(w.rankInfo.ScoreNum) {
		return
	}
	pointer := atomic.LoadPointer(&w.rank)
	topRank := (*[]*pb.RankUnitData)(pointer)
	pos, _ := SearchByScore(*topRank, len(*topRank), score)
	rank = uint32(pos)
	return
}

func (w *WholeRank) Stop() {
	if !w.running {
		return
	}
	w.running = false
	w.updateTick.Stop()
	w.close <- 1 //等待数据写入线程退出再close文件
	w.pstCache.Close()
}

func (w *WholeRank) Close() {
	w.Stop()
	os.Remove(w.cachePath)
}
