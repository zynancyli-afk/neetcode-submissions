func letterCombinations(digits string) []string {
	val := map[int][]string{
		2:[]string{"a","b","c"},
		3:[]string{"d","e","f"},
		4:[]string{"g","h","i"},
		5:[]string{"j","k","l"},
		6:[]string{"m","n","o"},
		7:[]string{"q","p","r","s"},
		8:[]string{"t","u","v"},
		9:[]string{"w","x","y","z"},
	}
	perms := []string{}
	var makePerm func(d string)
	curr := ""
	makePerm = func(d string){
		if len(curr)==len(digits){
			perms = append(perms, curr)
			return
		}
		c,_ := strconv.Atoi(d[0:1])
		for _,w := range val[c]{
			curr+=w
			makePerm(d[1:])
			curr=curr[:len(curr)-1]
		}
	}
	if len(digits)>0{makePerm(digits)}
	return perms
}
