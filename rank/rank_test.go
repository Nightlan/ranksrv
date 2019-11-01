package rank

import (
	"fmt"
	"math/rand"
	"ranksrv/proto/pb"
	"strconv"
	"testing"
	"time"
)

const (
	PartRankName = "part_rank"
)

var (
	rankMgr *Manager
)

func TestRank(t *testing.T) {
	if InitRankMgr(`D:\dcache\`, "10091") != nil {
		fmt.Println("Init failed")
		return
	}
	rankMgr = GetMgrInstance()
	partRank := &pb.RankInfo{
		Name:          PartRankName,
		Type:          pb.RANK_TYPE_WHOLE,
		TopNum:        10000,
		FlushInterval: 60,
		ScoreNum:      2,
		ScoreMinLimit: 0,
	}
	if err := rankMgr.RankInit(partRank); err != nil {
		fmt.Println("Init rank failed", err)
	}
	for i := 0; i < 10000000; i++ {
		uid := strconv.Itoa(i)
		data := &pb.RankUnitData{
			UniqueID:   uid,
			UpdateTime: uint32(time.Now().Unix()),
		}
		for j := 0; j < 2; j++ {
			data.Score = append(data.Score, uint64(rand.Intn(100000)))
		}
		if err := rankMgr.RankUpdate(PartRankName, data); err != nil {
			fmt.Println("update err")
		}
	}

	for i := 0; i < 1000; i++ {
		go update(i)
	}
	for {
		rankInfo, err := rankMgr.RankQuery(PartRankName, "1")
		if err != nil {
			fmt.Println("query rank failed", err)
		} else {
			fmt.Println("rank info", rankInfo)
		}
		topList, err := rankMgr.RankQueryTop(PartRankName, 100)
		if err != nil {
			fmt.Println("query top failed", err)
		} else {
			fmt.Println("top list", len(topList), topList)
		}
		score := []uint64{4, 2}
		rank, err := rankMgr.RankQueryByScore(PartRankName, score)
		if err != nil {
			fmt.Println("query top failed", err)
		} else {
			fmt.Println("rank", rank)
		}
		time.Sleep(time.Second)
	}
	if CloseRankMgr() != nil {
		fmt.Println("Close failed")
	}
}

func update(i int) {
	uid := strconv.Itoa(i)
	for {
		data := &pb.RankUnitData{
			UniqueID:   uid,
			UpdateTime: uint32(time.Now().Unix()),
		}
		for j := 0; j < 2; j++ {
			data.Score = append(data.Score, uint64(rand.Intn(100000)))
		}
		if err := rankMgr.RankUpdate(PartRankName, data); err != nil {
			fmt.Println("update err")
		}
		time.Sleep(time.Second)
	}
}
