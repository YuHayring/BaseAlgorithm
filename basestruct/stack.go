package basestruct

type Stack struct {
	Head *ListNode
}

func (this *Stack) push(val interface{}) {
	node := new(ListNode)
	node.Val = val
	this.Head.Next, node.Next = node, this.Head.Next
}

func (this *Stack) pull() interface{} {
	val := this.Head.Next.Val
	this.Head.Next = this.Head.Next.Next
	return val
}

func (this *Stack) Empty() bool {
	return this.Head.Next == nil
}

