func canFinish(numCourses int, prerequisites [][]int) bool {
	prereq := make([]int,10000)
	unblock := make(map[int][]int)
	for _,v := range prerequisites {
		list, _ := unblock[v[1]]
		list=append(list, v[0])
		unblock[v[1]]=list
		prereq[v[0]]++
	}
	canTake := []int{}
	for i := 0; i < numCourses; i++{
		if prereq[i] == 0{canTake=append(canTake, i)}
	}
	taken := 0
	for len(canTake)>0{
		taking := canTake[0]
		canTake=canTake[1:]
		taken++
		unblocklist,_:=unblock[taking]
		for _,v := range unblocklist {
			prereq[v]--
			if prereq[v]==0{
				canTake=append(canTake, v)
			}
		}
	}
	return taken >= numCourses
}
