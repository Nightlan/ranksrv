package core

import (
	"ranksrv/proto/pb"
)

type RegularArray struct {
	endIndex     int
	currentIndex int
	buf          []*pb.RankUnitData
}

func NewRegularArray(size int) *RegularArray {
	return &RegularArray{
		endIndex:     size - 1,
		currentIndex: -1,
		buf:          make([]*pb.RankUnitData, size),
	}
}

func (a *RegularArray) Size() int {
	return a.currentIndex + 1
}

func (a *RegularArray) Find(data *pb.RankUnitData) (int, bool) {
	return SearchByData(a.buf, a.currentIndex+1, data)
}

func (a *RegularArray) FindByScore(score []uint64) (int, bool) {
	return SearchByScore(a.buf, a.currentIndex+1, score)
}

func (a *RegularArray) GetTop(reqNum int) (top []*pb.RankUnitData) {
	if reqNum > a.currentIndex+1 {
		reqNum = a.currentIndex + 1
	}
	top = make([]*pb.RankUnitData, reqNum)
	copy(top, a.buf)
	return
}

func (a *RegularArray) Append(data *pb.RankUnitData) bool {
	if a.currentIndex == a.endIndex {
		return false
	}
	a.currentIndex++
	a.buf[a.currentIndex] = data
	return true
}

func (a *RegularArray) Update(oldPos, newPos int, data *pb.RankUnitData) {
	if oldPos == newPos {
		a.buf[newPos] = data
		return
	} else if oldPos > newPos {
		for oldPos > newPos {
			a.buf[oldPos] = a.buf[oldPos-1]
			oldPos--
		}
		a.buf[newPos] = data
	} else {
		curPos := newPos - 1
		for oldPos < curPos {
			a.buf[oldPos] = a.buf[oldPos+1]
			oldPos++
		}
		a.buf[curPos] = data
	}
}

func (a *RegularArray) Insert(pos int, data *pb.RankUnitData) (endData *pb.RankUnitData) {
	if pos > a.currentIndex {
		return
	}
	a.currentIndex++
	if a.currentIndex > a.endIndex {
		a.currentIndex = a.endIndex
		endData = a.buf[a.endIndex]
	}
	size := a.currentIndex
	for pos < size {
		a.buf[size] = a.buf[size-1]
		size--
	}
	a.buf[pos] = data
	return
}

func (a *RegularArray) Delete(pos int) {
	if pos > a.currentIndex {
		return
	}
	for pos < a.currentIndex {
		a.buf[pos] = a.buf[pos+1]
		pos++
	}
	a.buf[a.currentIndex] = nil
	a.currentIndex--
}
