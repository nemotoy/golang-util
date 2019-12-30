package slice

// UInt returns a unique slice of int from the given slice of int.
func UInt(ii []int) []int {
	m := make(map[int]struct{})
	ids := make([]int, 0, len(ii))
	for _, i := range ii {
		if _, ok := m[i]; !ok {
			m[i] = struct{}{}
			ids = append(ids, i)
		}
	}
	return ids
}
