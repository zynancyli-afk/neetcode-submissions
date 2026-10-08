type WordDictionary struct {
    char rune
	children map[rune]*WordDictionary
	isEnd bool
}

func Constructor() WordDictionary {
    return WordDictionary{
		children : make(map[rune]*WordDictionary),
		isEnd: false,
	}
}

func (this *WordDictionary) AddWord(word string)  {
    tree := this
	for _,v := range word {
		if tree == nil {
			return
		}
		child, ok := tree.children[v]
		if !ok {
			c := Constructor()
			child = &c
			child.char = v
			tree.children[v]=child
		}
		tree = child
	}
	tree.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    trees := []*WordDictionary{this}
	for _,v := range word {
		newTree := []*WordDictionary{}
		for _,t := range trees {
			if t == nil {
				continue
			}
			if v == '.'{
				for _,c := range t.children {
					newTree = append(newTree, c)
				}
			}else {
				v, ok := t.children[v]
				if ok {
					newTree = append(newTree, v)
				}
			}

		}
		trees = newTree
		if len(newTree)== 0 {
			break
		}
	}
	for _,t := range trees {
		if t.isEnd {
			return true
		}
	}
	return false
}
