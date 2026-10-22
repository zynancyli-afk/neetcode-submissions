import "slices"

func insert(intervals [][]int, newInterval []int) [][]int {
	insertIndex := -1
    for i,v := range intervals {
		if newInterval[0] <= v[1]{
			insertIndex = i
			break
		}
	}
	if insertIndex == -1 {
		return append(intervals, newInterval)
	}
	next := insertIndex
	for next < len(intervals) && doesOverlap(newInterval, intervals[next]){
		newInterval = overlap(newInterval, intervals[next])
		next++
	}
	return slices.Concat(intervals[:insertIndex], [][]int{newInterval}, intervals[next:])
	
}

func doesOverlap(interval1, interval2 []int) bool {
	return interval1[0] <= interval2[1] && interval2[0] <= interval1[1]
}

func overlap(interval1, interval2 []int)[]int{
	start := min(interval1[0], interval2[0])
	end := max(interval1[1], interval2[1])
	return []int{start, end}
}
