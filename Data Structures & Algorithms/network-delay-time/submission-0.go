import "slices"

func networkDelayTime(times [][]int, n int, k int) int {
	timemap := make(map[int][]int)
	for i,v := range times {
		tm,_ := timemap[v[0]]
		tm = append(slices.Clone(tm), i)
		timemap[v[0]]=tm
	}
    
	reach:=make([]int, n+1)
	for i,_ :=range reach {reach[i]=-1}
	applied := []int{k}
	reach[k]=0
	for len(applied)>0{
		tm := timemap[applied[0]]
		applied = applied[1:]
		for _,v := range tm {
			v := times[v]
			node := v[1]
			time := reach[v[0]] + v[2]
			if reach[node]==-1 || time<reach[node]{
				reach[node]=time
				applied=append(slices.Clone(applied), node)
			}
		}
	}
	max := 0
	for i,v := range reach {
		if i == 0 {continue}
		if v == -1 {return -1}
		if v > max {max = v}
	}
	return max
}
