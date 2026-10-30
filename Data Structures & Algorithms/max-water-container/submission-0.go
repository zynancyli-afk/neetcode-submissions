func maxArea(heights []int) int {
	max := 0
	l,r := 0,len(heights)-1
	for l < r {
		area := min(heights[l],heights[r])*(r-l)
		if area > max {
			max = area
		}
		if heights[l] < heights[r]{
			l++
		}else {r--}
	}
	return max
}
