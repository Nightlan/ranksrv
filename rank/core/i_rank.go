package core

import (
	"ranksrv/proto/pb"
)

type IRank interface {
	Start(cacheDir string) (err error)
	RankInfo() (rankInfo *pb.RankInfo)
	Update(data *pb.RankUnitData)
	Delete(uniqueID string)
	QueryRank(uniqueID string) (rankInfo *pb.UnitRankInfo, err error)
	QueryTop(reqNum uint32) (topList []*pb.RankUnitData, err error)
	QueryByScore(score []uint64) (rank uint32, err error)
	Stop()
	Close()
}
