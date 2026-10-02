import "slices"

func lastStoneWeight(stones []int) int {
	slices.Sort(stones)
	for len(stones) > 1 {
		s1 := stones[len(stones)-1]
		s2 := stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if s1 != s2 {
			stones = append(stones,max(s1-s2, s2-s1))
			slices.Sort(stones)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
