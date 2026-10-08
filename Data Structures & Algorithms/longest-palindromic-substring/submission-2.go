func longestPalindrome(s string) string {
	longest := ""
	chars := []rune(s)
    for i,v := range chars {
		lower,upper:=i,i
		for lower >= 0 && upper < len(s) && chars[lower]==chars[upper]{
			substring := s[lower:upper+1]
			if len(substring)>len(longest){
				longest = substring
			}
			lower--
			upper++
		}
		if i > 0 && v == chars[i-1] {
			lower = i-1
			upper = i
			for lower >= 0 && upper < len(s) && chars[lower]==chars[upper]{
				substring := s[lower:upper+1]
				if len(substring)>len(longest){
					longest = substring
				}
				lower--
				upper++
			}
		}
	}
	return longest
}
