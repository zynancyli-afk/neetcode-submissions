func coinChange(coins []int, amount int) int {
	possibleAmount := make(map[int]int)
	for _,v := range coins {
		possibleAmount[v]=1
	}
	possibleAmount[0]=0

	var subCoinChange func(a int)int
	subCoinChange = func(a int)int{
		if val, ok := possibleAmount[a]; ok {
			return val
		}
		minCoins := -1
		for _,v := range coins {
			if a > v {
				subCoins := subCoinChange(a-v) 
				if subCoins != -1 && (minCoins == -1 || minCoins > subCoins) {
					minCoins = subCoins
				}
			}
		}
		if minCoins != -1 {
			minCoins++
		}
		possibleAmount[a]=minCoins
		//fmt.Println(a,minCoins)
		return minCoins
	}

	return subCoinChange(amount)
}




