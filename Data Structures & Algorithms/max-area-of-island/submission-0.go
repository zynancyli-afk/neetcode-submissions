func maxAreaOfIsland(grid [][]int) int {
	visited := make([][]bool,len(grid))
	for i,_ := range visited{
		visited[i]=make([]bool, len(grid[0]))
	}


	direction := [][]int{
		[]int{0,-1},
		[]int{0,1},
		[]int{1,0},
		[]int{-1,0},
	}

	var visit func(x,y int)int
	visit = func(x,y int)int{
		visited[x][y]=true
		count:=1
		for _,d := range direction {
			xx := x+d[0]
			yy := y+d[1]
			if xx < 0 || xx >= len(grid) || yy < 0 || yy >= len(grid[0]) || grid[xx][yy]==0{
				continue
			}
			if !visited[xx][yy] {
				count += visit(xx,yy)
			}
		}
		return count
	}

    maxA := 0
	for i,v := range grid {
		for j,k := range v {
			if k == 0 {continue}
			if !visited[i][j]{
				count := visit(i,j)
				if count > maxA {
					maxA = count
				}
			}
		}
	}
	return maxA
}
