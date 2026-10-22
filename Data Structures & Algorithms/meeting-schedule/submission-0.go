/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

 import "slices"

func canAttendMeetings(intervals []Interval) bool {
	slices.SortFunc(intervals, func(a,b Interval) int {
		return a.start - b.start 
	})
	lastStart := 0
	for _,v :=range intervals {
		if v.start < lastStart {
			return false
		}
		lastStart = v.end
	}
	return true
}
