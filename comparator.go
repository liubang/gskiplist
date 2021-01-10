package gskiplist

type Comparator interface {
	Compare(a, b interface{}) int
}

type IntComparator struct {
	Comparator
}

func (ic *IntComparator) Compare(a, b interface{}) int {
	return a.(int) - b.(int)
}

func NewIntComparator() *IntComparator {
	return &IntComparator{}
}
