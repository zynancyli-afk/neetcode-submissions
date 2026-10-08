func countSubstrings(s string) int {   
	count := 0
	chars := []rune(s)
    for i,v := range chars {
		lower,upper:=i,i
		for lower >= 0 && upper < len(s) && chars[lower]==chars[upper]{
			lower--
			upper++
			count++
		}
		if i > 0 && v == chars[i-1] {
			lower = i-1
			upper = i
			for lower >= 0 && upper < len(s) && chars[lower]==chars[upper]{
				lower--
				upper++
				count++
			}
		}
	}
	return count
}
