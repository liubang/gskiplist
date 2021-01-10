package gskiplist

import "sync"

type Node struct {
	key   interface{}
	next  []*Node
	mutex sync.RWMutex
}

func newNode(key interface{}, height int) *Node {
	return &Node{
		key:  key,
		next: make([]*Node, height),
	}
}

func (node *Node) Next(n int) *Node {
	node.mutex.RLock()
	defer node.mutex.RUnlock()
	return node.next[n]
}

func (node *Node) SetNext(n int, x *Node) {
	node.mutex.Lock()
	defer node.mutex.Unlock()
	node.next[n] = x
}
