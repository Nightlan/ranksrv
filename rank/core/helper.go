package core

type CompareResult int

const (
	Equal CompareResult = 1 + iota
	Greater
	Less
)

func HalfSearch(n int, f func(int) CompareResult) (int, bool) {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		rst := f(h)
		if rst == Greater {
			j = h
		} else if rst == Less {
			i = h + 1
		} else {
			return h, true
		}
	}
	return i, false
}
