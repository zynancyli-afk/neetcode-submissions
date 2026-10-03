func subsets(nums []int) [][]int {
	sets := [][]int{
		[]int{},
	}
	for _,v := range nums{
		n := len(sets)
		for i := 0; i < n; i++ {
			next := make([]int, len(sets[i]), len(sets[i])+1)
			copy(next, sets[i])
			next = append(next, v)
			sets = append(sets, next)
		}
	}
	return sets
}
