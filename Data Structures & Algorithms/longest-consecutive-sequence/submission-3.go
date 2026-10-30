func longestConsecutive(nums []int) int {
	have := make(map[int]bool, len(nums))
	for _, v := range nums {
		have[v] = true
	}
	longest := 0
	for v := range have { // iterating the set also skips duplicates
		if have[v-1] {
			continue // not the start of a sequence
		}
		count := 0
		for have[v+count] {
			count++
		}
		if count > longest {
			longest = count
		}
	}
	return longest
}