import "slices"

func partition(s string) [][]string {
	possible := [][]string{}
	set := []string{}
	var part func(start int)
	part = func(start int){
		if start == len(s) {
			possible = append(possible,slices.Clone(set))
		}
		for end:= start+1;end <= len(s);end++{
			str := s[start:end]
			if isPalindrom(str){
				set = append(set, str)
				part(end)
				set = set[:len(set)-1]
			}
		}
	}
	part(0)
	return possible
}

func isPalindrom(s string)bool{
	for i:=0;i<len(s)/2;i++{
		if s[i]!=s[len(s)-1-i]{return false}
	}
	return true
}

