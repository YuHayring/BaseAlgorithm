package main


//并查集
type UnionFind struct {
	size int
	pre  []int

}

func ConstructorOfUnionFind(cap int) *UnionFind {
	set := new(UnionFind)
	set.pre = make([]int, cap)
	for i,_ := range set.pre {
		set.pre[i] = i
	}
	set.size = cap
	return set
}

//待优化
func (this *UnionFind) union(a int, b int) {
	if this.find(a) == this.find(b) {
		return
	}
	this.pre[a] = b
	this.size--
}

func (this *UnionFind) find(me int) int {
	for me != this.pre[me] {
		me, this.pre[me] = this.pre[me], this.pre[this.pre[me]]
	}
	return me
}
