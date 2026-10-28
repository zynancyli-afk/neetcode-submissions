/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

 import "slices"

func minMeetingRooms(intervals []Interval) int {
	concurrent := 0
	start := []int{}
	end := []int{}
	for _,v := range intervals {
		start = append(start, v.start)
		end = append(end, v.end)
	}
	slices.Sort(start)
	slices.Sort(end)

	s,e := 0,0
	curr := 0
	for s < len(intervals){
		if start[s]<end[e]{
			curr++
			s++
		}else {
			curr--
			e++
		}
		if curr > concurrent {concurrent = curr}
	}
	return concurrent
}


