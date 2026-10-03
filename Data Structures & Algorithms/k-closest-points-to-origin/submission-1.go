import "slices"

func kClosest(points [][]int, k int) [][]int {
	closestPoints := [][]int{}
	max := float64(0)
	for _,v := range points {
		distance := math.Hypot(float64(v[0]),float64(v[1]))
		if len(closestPoints) < k {
			closestPoints = append(closestPoints, v)
			if distance > max {
				max = distance
			}
		}else if distance < max {
			closestPoints = append(closestPoints, v)
			evictpos := 0
			newmax := distance
			for j,w := range closestPoints {
				dis := math.Hypot(float64(w[0]),float64(w[1]))
				if dis > newmax && dis != max {
					newmax = dis
				}else if dis == max {
					evictpos = j
				}
			}
			max = newmax
			closestPoints = slices.Delete(closestPoints,evictpos,evictpos+1)
		}
	}
	return closestPoints
}
