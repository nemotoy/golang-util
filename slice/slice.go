package slice

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
