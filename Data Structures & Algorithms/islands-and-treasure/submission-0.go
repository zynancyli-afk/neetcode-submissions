func islandsAndTreasure(grid [][]int) {
    directions := [][]int{
		{0,1},{0,-1},{1,0},{-1,0},
	}
	startingPos := [][]int{}
	for i,v := range grid {
		for j,k := range v {
			if k==0{
				startingPos = append(startingPos, []int{i,j})
			}
		}
	}
	for len(startingPos) > 0 {
		newPos := [][]int{}
		for _,v := range startingPos {
			for _,w := range directions {
				x, y := v[0]+w[0], v[1]+w[1]
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y]==-1 {
					continue
				}
				if grid[x][y] > grid[v[0]][v[1]]+1{
					grid[x][y] = grid[v[0]][v[1]]+1
					newPos=append(newPos, []int{x,y})
				}
			}
		}
		startingPos=newPos
	}

}
