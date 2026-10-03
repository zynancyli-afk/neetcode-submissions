type sumArr struct {
	arr   []int
	sum   int
	start int // index in nums of the last number used
}

func combinationSum(nums []int, target int) [][]int {
	unmapped := []sumArr{{arr: []int{}, sum: 0, start: 0}}
	combos := [][]int{}
	for len(unmapped) > 0 {
		newunmapped := []sumArr{}
		for _, v := range unmapped {
			for i := v.start; i < len(nums); i++ {
				w := nums[i]
				newsum := v.sum + w
				if newsum > target {
					continue
				}
				combo := make([]int, len(v.arr), len(v.arr)+1)
				copy(combo, v.arr)
				combo = append(combo, w)
				if newsum == target {
					combos = append(combos, combo)
				} else {
					newunmapped = append(newunmapped, sumArr{combo, newsum, i})
				}
			}
		}
		unmapped = newunmapped
	}
	return combos
}