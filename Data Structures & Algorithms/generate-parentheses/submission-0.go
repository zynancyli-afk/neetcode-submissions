func generateParenthesis(n int) []string {
	collection := []string{}
	var generateWithPrefix func(pre string, open int, close int)
	generateWithPrefix = func(pre string, open int, close int){
		if len(pre) == n*2 && open == close {
			collection = append(collection, pre)
			return
		}
		//close
		if open > close {
			generateWithPrefix(pre+")",open, close+1)
		}
		//open
		if open < n {
			generateWithPrefix(pre+"(", open+1, close)
		}
	}
	generateWithPrefix("",0, 0)
	return collection
}
