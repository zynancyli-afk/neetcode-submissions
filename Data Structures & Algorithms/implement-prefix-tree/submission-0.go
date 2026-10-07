type PrefixTree struct {
	val rune
	children map[rune]*PrefixTree
	isEnd bool
}

func Constructor() PrefixTree {
    return PrefixTree{
		children:make(map[rune]*PrefixTree),
		isEnd:false,
	} 
}

func (this *PrefixTree) Insert(word string) {
	if this == nil {
		return
	}
	tree := this
	for _,v := range word {
		childtree,ok := tree.children[v]
		if !ok {
			childtree = &PrefixTree{
				val:v,
				children: make(map[rune]*PrefixTree),
			}
			tree.children[v]=childtree
		}
		tree = childtree
	}
	tree.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	tree := this
	for _,v := range word {
		if tree == nil {
			return false
		}
		val,ok := tree.children[v]; 
		if !ok {
			return false
		}
		tree = val
	}
	return tree.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	tree := this
	for _,v := range prefix {
		if tree == nil {
			return false
		}
		val, ok := tree.children[v]
		if !ok {
			return false
		}
		tree = val
	}
	return true
}
