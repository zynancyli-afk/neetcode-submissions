func isHappy(n int) bool {
    seen := make(map[int]bool)
	for n > 0 {
		if n == 1 {
			return true
		}
		if yes, ok := seen[n]; ok && yes {
			return false
		}
		seen[n]=true
		num := 0
		for n > 0 {
			r := n%10
			num+=r*r
			n = n/10
		}
		n = num
	}
	return false
}
