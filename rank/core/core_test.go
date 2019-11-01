package core

import (
	"fmt"
	"math/rand"
	"ranksrv/proto/pb"
	"sort"
	"strconv"
	"testing"
)

func TestCore(t *testing.T) {
	intArray := make([]int, 0, 10)
	for i := 10; i > 0; i-- {
		intArray = append(intArray, i*2)
	}
	intKey := 1
	pos1, ok1 := HalfSearch(len(intArray), func(i int) CompareResult {
		if intArray[i] > intKey {
			return Less
		} else if intArray[i] < intKey {
			return Greater
		} else {
			return Equal
		}
	})
	fmt.Println("Test int")
	fmt.Println(intArray)
	fmt.Println(pos1, ok1)
	fmt.Println("Test int end")
	array := make([]*pb.RankUnitData, 0, 100)
	for i := 0; i < 100; i++ {
		data := &pb.RankUnitData{
			UniqueID:   strconv.Itoa(rand.Intn(5)),
			UpdateTime: uint32(rand.Intn(1)),
		}
		data.Score = append(data.Score, uint64(rand.Intn(5)))
		for j := 0; j < 2; j++ {
			data.Score = append(data.Score, uint64(rand.Intn(5)))
		}
		array = append(array, data)
	}
	sort.Slice(array, func(i, j int) bool {
		return UnitDataLess(array[i], array[j])
	})
	for i, v := range array {
		fmt.Println(v, i)
	}
	key := &pb.RankUnitData{
		UniqueID:   "1",
		UpdateTime: 1,
	}
	key.Score = append(key.Score, 4)
	key.Score = append(key.Score, 0)
	key.Score = append(key.Score, 3)
	fmt.Println(key)
	pos, ok := HalfSearch(len(array), func(i int) CompareResult {
		return UnitDataCompare(array[i], key)
	})
	fmt.Println(pos, ok)
	score := make([]uint64, 0)
	score = append(score, 4)
	score = append(score, 0)
	score = append(score, 2)
	fmt.Println(score)
	pos, ok = SearchByScore(array, len(array), score)
	fmt.Println(pos, ok)
}
