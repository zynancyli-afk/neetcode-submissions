import "slices"


func permute(nums []int) [][]int {
	perms := [][]int{
		{},
	}
	for _,v := range nums {
		newPerms := [][]int{}
		for _,w := range perms {
			for i := 0; i <= len(w); i++ {
				p := slices.Concat(w[:i], []int{v},w[i:])
				newPerms = append(newPerms, p)
			}
		}
		perms = newPerms
	}
	return perms
}
