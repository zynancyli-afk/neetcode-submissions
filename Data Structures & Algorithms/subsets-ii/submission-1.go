import "slices"

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	perms := [][]int{
		[]int{},
	}

	var gensubset func(subset[]int, index int)
	gensubset = func(subset[]int, index int){
		last := -22
		for i := index; i < len(nums); i++ {
			if nums[i]==last {
				continue
			}
			newsubset := append(slices.Clone(subset), nums[i])
			perms = append(perms, newsubset)
			gensubset(newsubset, i+1)
			last = nums[i]
		}
	}
	gensubset([]int{},0)
	return perms
}
