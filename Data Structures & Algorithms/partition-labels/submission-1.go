func partitionLabels(s string) []int {
	fmt.Println(s)
	minMax := [26][2]int{}
	for i,v := range s {
		if minMax[v-'a'][0]==0{
			minMax[v-'a'][0]=i
		}
		minMax[v-'a'][1]=i
	}
	max := 0
	lastmax := -1
	splits := []int{}
	for i,v := range s {
		bounds := minMax[v-'a']
		if bounds[1]> max { max = bounds[1] }
		if i == max {
			splits = append(splits,i-lastmax)
			lastmax = max
		}
	}
	return splits
}
