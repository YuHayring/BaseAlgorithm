package basestruct

type Comparable interface {
	CompareTo(elem *Comparable) int
}
