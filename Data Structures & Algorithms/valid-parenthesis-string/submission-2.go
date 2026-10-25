func checkValidString(s string) bool {
	openIndex := []int{}
	optionalIndex := []int{}
	for i,v := range s {
		switch v {
			case '(':
				openIndex = append(openIndex, i)
			case '*':
				optionalIndex = append(optionalIndex, i)
			case ')':
				if len(openIndex) > 0 {
					openIndex = openIndex[:len(openIndex)-1]
				}else if len(optionalIndex) > 0 {
					optionalIndex=optionalIndex[:len(optionalIndex)-1]
				}else {return false}
		}
	} 
	for i:=len(openIndex)-1; i >= 0; i-- {
		if len(optionalIndex) <= 0 {
			return false
		}
		if optionalIndex[len(optionalIndex)-1] > openIndex[i] {
			optionalIndex = optionalIndex[:len(optionalIndex)-1]
		}else {
			return false
		}
	}
	return true
}
