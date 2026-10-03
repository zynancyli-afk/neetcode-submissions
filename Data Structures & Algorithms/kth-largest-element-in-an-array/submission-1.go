import "slices"

func findKthLargest(nums []int, k int) int {
	klargest := []int{}
	for _,v := range nums {
		if len(klargest) < k {
			klargest = append(klargest,v)
			if len(klargest) == k {
				slices.Sort(klargest)
			}
		}else if v > klargest[0] {
			klargest = append(klargest,v)
			slices.Sort(klargest)
			klargest = klargest[1:]
		}
	}
	return klargest[0]
}
