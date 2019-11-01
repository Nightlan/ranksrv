package core

import (
	"ranksrv/proto/pb"
)

type CacheUnitData struct {
	*pb.RankUnitData
	Update bool
}

func UnitDataCompare(i, j *pb.RankUnitData) CompareResult {
	scoreNum := len(i.Score)
	for d := 0; d < scoreNum; d++ {
		if i.Score[d] < j.Score[d] {
			return Greater
		} else if i.Score[d] == j.Score[d] {
			continue
		} else {
			return Less
		}
	}
	if i.UpdateTime < j.UpdateTime {
		return Less
	} else if i.UpdateTime > j.UpdateTime {
		return Greater
	} else {
		if i.UniqueID < j.UniqueID {
			return Less
		} else if i.UniqueID > j.UniqueID {
			return Greater
		}
	}
	return Equal
}

func UnitDataLess(i, j *pb.RankUnitData) bool {
	compareRst := UnitDataCompare(i, j)
	if compareRst == Less {
		return true
	}
	return false
}

func SearchByData(array []*pb.RankUnitData, size int, data *pb.RankUnitData) (int, bool) {
	if size > len(array) {
		size = len(array)
	}
	return HalfSearch(size, func(i int) CompareResult {
		return UnitDataCompare(array[i], data)
	})
}

func SearchByScore(array []*pb.RankUnitData, size int, score []uint64) (int, bool) {
	if size > len(array) {
		size = len(array)
	}
	return HalfSearch(size, func(i int) CompareResult {
		scoreNum := len(score)
		for d := 0; d < scoreNum; d++ {
			if array[i].Score[d] < score[d] {
				return Greater
			} else if array[i].Score[d] == score[d] {
				continue
			} else {
				return Less
			}
		}
		return Equal
	})
}
