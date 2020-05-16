package basestruct

var BASE byte = 'a'

type TrieTree struct {
	next [26]*TrieTree
	end bool
}

func BuildTrieTree() *TrieTree {
	return &TrieTree{[26]*TrieTree{}, false}
}

func  (this *TrieTree) Insert(word string) {
	str := []byte(word)
	p := this
	for _,char := range str {
		if p.next[char - BASE] == nil {
			p.next[char - BASE] = BuildTrieTree()
		}
		p = p.next[char - BASE]
	}
	p.end = true
}

func (this *TrieTree) Search(word string) bool {
	str := []byte(word)
	p := this
	for _,char := range str {
		if p.next[char - BASE] == nil {
			return false
		} else {
			p = p.next[char - BASE]
		}
	}
	return p.end
}

func (this *TrieTree) StartsWith(prefix string) bool {
	str := []byte(prefix)
	p := this
	for _,char := range str {
		if p.next[char - BASE] == nil {
			return false
		} else {
			p = p.next[char - BASE]
		}
	}
	return true
}