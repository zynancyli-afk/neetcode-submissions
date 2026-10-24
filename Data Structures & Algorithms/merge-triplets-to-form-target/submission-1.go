func mergeTriplets(triplets [][]int, target []int) bool {
	one, two, three := 0,0,0
    for _,v := range triplets {
		if v[0]<=target[0] && v[1]<= target[1] && v[2]<=target[2] {
			one = max(v[0],one)
			two = max(v[1], two)
			three = max(v[2], three)
		}
	}
	return one == target[0] && two == target[1] && three == target[2]
}
