import "slices"

func merge(intervals [][]int) [][]int {
    merged := [][]int{}
	slices.SortFunc(intervals, func(a,b []int)int{
		if a[0]== b[0]{
			return a[1]-b[1]
		}
		return a[0]-b[0]
	})
	last := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= last[1] {
			last[1]=max(intervals[i][1],last[1])
		}else {
			merged = append(merged, last)
			last = intervals[i]
		}
	}
	return append(merged, last)
}
