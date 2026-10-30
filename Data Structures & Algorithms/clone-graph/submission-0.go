/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	cloned := make(map[*Node]*Node)
	var clone func(node *Node)*Node
	clone = func(node *Node)*Node{
		if node == nil {
			return nil
		}
		if val, ok := cloned[node];ok {
			return val
		}
		copy := &Node{
			Val: node.Val,
		}
		cloned[node]=copy
		neighbors := []*Node{}
		for _,v := range node.Neighbors {
			neighbors = append(neighbors, clone(v))
		}
		copy.Neighbors=neighbors
		return copy
	}
	return clone(node)
}
