package basestruct

type Queue struct {
	Last *ListNode
	Head *ListNode
	len int
}

func ConstructorOfQueue() *Queue {
	this := new(Queue)
	this.Head = new(ListNode)
	this.Last = this.Head
	return this
}

func BuildQueueByIntArray(arr []int) *Queue {
	this := new(Queue)
	this.Head = new(ListNode)
	this.Last = this.Head
	for _, val := range arr {
		this.Push(val)
	}
	return this
}



func (this *Queue) Push(val interface{}) {
	this.Last.Next = new(ListNode)
	this.Last.Next.Val = val
	this.Last = this.Last.Next
	this.len++
}

func (this *Queue) Pull() interface{} {
	val := this.Head.Next.Val
	this.Head.Next = this.Head.Next.Next
	if this.Head.Next == nil {
		this.Last = this.Head
	}
	this.len--
	return val

}

func (this *Queue) Empty() bool {
	return this.Head.Next == nil
}

func (this *Queue) Len() int {
	return this.len
}