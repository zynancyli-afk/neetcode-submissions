func isValidSudoku(board [][]byte) bool {
	col := make([]bool,100)   //col,number
	grid := make([]bool, 1000) //row,col,number

	for i,v := range board {
		row := make([]bool,10)
		for j,w := range v {
			if w == '.' {continue}
			ww := int(w-'0')
			if row[ww] || col[j*10+ww] || grid[i/3*100+j/3*10+ww] {return false}
			row[ww]=true
			col[j*10+ww]=true
			grid[i/3*100+j/3*10+ww]=true
		}
	}
	return true
}
