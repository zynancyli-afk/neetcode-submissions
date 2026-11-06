func orangesRotting(grid [][]int) int {
	rotten := [][]int{}
	fresh := 0
    for x,v := range grid{
		for y,w := range v {
			if w == 2 {rotten = append(rotten,[]int{x,y})}
			if w == 1 {fresh++}
		}
	}
	if fresh == 0 {
		return 0
	}
	minute := -1
	direction := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
	for len(rotten)> 0 {
		fmt.Println(rotten)
		newlyRotten := [][]int{}
		for _,v := range rotten {
			for _,w := range direction {
				x := v[0]+w[0]
				y := v[1]+w[1]
				if x >= 0 && x < len(grid)&& y>=0 && y < len(grid[0]) && grid[x][y]==1{
					grid[x][y]=2
					newlyRotten=append(newlyRotten, []int{x,y})
					fresh--
				}
			}
		}
		minute++
		rotten=newlyRotten
	}
	if fresh > 0 {return -1}
	return minute

}
