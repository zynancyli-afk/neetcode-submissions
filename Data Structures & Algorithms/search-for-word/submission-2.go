func exist(board [][]byte, word string) bool {
	direction := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
	visited := make([][]bool, len(board)+1)
	for i:=0;i<len(board[0]);i++{visited[i]=make([]bool, len(board[0])+1)}

	var backtrack func(x,y,i int)bool
	backtrack = func(x,y,i int)bool{
		if board[x][y]==word[i]{
			if i == len(word)-1{return true}
			visited[x][y]=true
			for _,d := range direction {
				xx := x+d[0]
				yy := y+d[1]
				if xx<0 || xx >= len(board) || yy<0 || yy >= len(board[0]) || visited[xx][yy]{continue}
				if backtrack(xx,yy,i+1){return true}
			}
			visited[x][y]=false
		}
		return false
	}
	for x,v := range board {
		for y,_ := range v {
			if backtrack(x,y,0){
				return true
			}
		}
	}

	return false
}
