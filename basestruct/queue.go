package basestruct

type Queue struct {
	Last *ListNode
	Head *ListNode
}

func ConstructorOfQueue() *Queue {
	this := new(Queue)
	this.Head = new(ListNode)
	this.Last = this.Head
	return this
}


func (this *Queue) Push(val interface{}) {
	this.Last.Next = new(ListNode)
	this.Last.Next.Val = val
	this.Last = this.Last.Next
}

func (this *Queue) Pull() interface{} {
	val := this.Head.Next.Val
	this.Head.Next = this.Head.Next.Next
	if this.Head.Next == nil {
		this.Last = this.Head
	}
	return val

}

func (this *Queue) Empty() bool {
	return this.Head.Next == nil
}
