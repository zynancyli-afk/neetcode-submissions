func myPow(x float64, n int) float64 {
	vals := make(map[int]float64)
	vals[0]=1
	vals[1]=x
	
	var find func(n int)float64
	find = func(n int)float64{
		if v,ok :=vals[n];ok{return v}
		half := find(n/2)
		half*=half
		if n%2 == 1{half*=x}
		vals[n]=half
		return half
	}
	if n < 0 {return 1/find(n*-1)}else{return find(n)}
}
