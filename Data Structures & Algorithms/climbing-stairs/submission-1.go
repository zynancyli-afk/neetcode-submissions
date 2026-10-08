func climbStairs(n int) int {
	if n == 0 {
		return 0
	}else if n == 1 {
		return 1
	}
	val := make([]int,n+1)
	val[1]=1
	val[2]=2
	for i:=3; i <= n; i++ {
		val[i]=val[i-1]+val[i-2]
	}
	return val[n]
}



/*
0 []
1 [1]
2 [1,1][2]
3 [1,1,1][1,2][2,1]
4 [1,1,1,1],
*/