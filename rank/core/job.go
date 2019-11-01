package core

import (
	"ranksrv/proto/pb"
)

type UpdateJob struct {
	Data *pb.RankUnitData
}

type DeleteJob struct {
	UniqueID string
}

type QueryRankJob struct {
	UniqueID   string
	RspChannel chan *pb.UnitRankInfo
}

type QueryTopJob struct {
	reqNum     int
	RspChannel chan []*pb.RankUnitData
}

type QueryByScoreJob struct {
	Score      []uint64
	RspChannel chan int
}
