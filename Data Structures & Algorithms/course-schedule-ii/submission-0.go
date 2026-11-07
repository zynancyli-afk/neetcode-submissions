func findOrder(numCourses int, prerequisites [][]int) []int {
	indegree := make([]int, 1000)
	unblock := make(map[int][]int)
	for _,v := range prerequisites {
		indegree[v[0]]++
		list, _ := unblock[v[1]]
		unblock[v[1]]=append(list,v[0])
	}
	justTaken := []int{}
	for i:=0; i < numCourses;i++ {
		if indegree[i]==0{
			justTaken=append(justTaken, i)
		}
	}
	order := []int{}
	for len(justTaken)>0{
		list,_ := unblock[justTaken[0]]
		order=append(order,justTaken[0])
		justTaken=justTaken[1:]
		for _,v := range list {
			indegree[v]--
			if indegree[v]==0{justTaken=append(justTaken, v)}
		}
	}
	if len(order)==numCourses {return order}
	return []int{}
}
