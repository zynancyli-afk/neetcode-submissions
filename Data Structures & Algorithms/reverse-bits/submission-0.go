func reverseBits(n int) int {
	result := 0
	for i:=0;i<32;i++{
		bit := (n>>i) & 1
		result |= bit<<(31-i)
	}
	return result
}
