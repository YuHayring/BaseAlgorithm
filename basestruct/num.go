package basestruct

type Integer int

func (this Integer) CompareTo(integer *Comparable) int {
	return int(this) - int((*integer).(Integer))
}
