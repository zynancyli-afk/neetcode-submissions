func numIslands(grid [][]byte) int {
	xlen := len(grid)
	ylen := len(grid[0])
	visited := make([][]bool, xlen)
	for i := range visited {
		visited[i] = make([]bool, ylen)
	}
	islands := 0
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			if grid[x][y] != '1' || visited[x][y] {
				continue
			}
			islands++
			visited[x][y] = true
			queue := [][2]int{{x, y}}
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				for _, d := range dirs {
					nx, ny := cur[0]+d[0], cur[1]+d[1]
					if nx < 0 || nx >= xlen || ny < 0 || ny >= ylen {
						continue
					}
					if grid[nx][ny] == '1' && !visited[nx][ny] {
						visited[nx][ny] = true
						queue = append(queue, [2]int{nx, ny})
					}
				}
			}
		}
	}
	return islands
}
