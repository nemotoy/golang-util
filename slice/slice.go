package slice

import "sort"

// UInt returns a unique slice of int from the given slice of int.
func UInt(raw []int) []int {
	m := make(map[int]bool, len(raw))
	ii := make([]int, 0, len(raw))
	for _, i := range raw {
		if _, ok := m[i]; !ok {
			m[i] = true
			ii = append(ii, i)
		}
	}
	return ii
}

// UIntASC returns sorted a unique slice of int in ascending order from the given slice of int.
func UIntASC(raw []int) []int {
	ii := UInt(raw)
	sort.Slice(ii, func(i, j int) bool {
		return ii[i] < ii[j]
	})
	return ii
}
