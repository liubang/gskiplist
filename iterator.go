package gskiplist

type Iterator struct {
	list *SkipList
	node *Node
}

func NewIterator(list *SkipList) *Iterator {
	return &Iterator{
		list: list,
		node: nil,
	}
}

func (iterator *Iterator) Valid() bool {
	return iterator.node != nil
}

func (iterator *Iterator) Key() interface{} {
	return iterator.node.key
}

func (iterator *Iterator) Next() {
	iterator.node = iterator.node.Next(0)
}

func (iterator *Iterator) Prev() {
	iterator.node = iterator.list.FindLessThan(iterator.node.key)
	if iterator.node == iterator.list.head {
		iterator.node = nil
	}
}

func (iterator *Iterator) Seek(key interface{}) {
	_, node := iterator.list.FindGreaterOrEqual(key)
	iterator.node = node
}

func (iterator *Iterator) SeekToFirst() {
	iterator.node = iterator.list.head.Next(0)
}

func (iterator *Iterator) SeekToLast() {
	iterator.node = iterator.list.FindLast()
	if iterator.node == iterator.list.head {
		iterator.node = nil
	}
}
