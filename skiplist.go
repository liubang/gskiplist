package gskiplist

import (
	"crypto/rand"
	"math/big"
	"sync/atomic"
)

const (
	kMaxHeight = 12
	kBranching = 4
)

type SkipList struct {
	maxHeight int64
	head      *Node
	compare   Comparator
}

func NewSkipList(comp Comparator) *SkipList {
	return &SkipList{
		maxHeight: 1,
		head:      newNode(nil, kMaxHeight),
		compare:   comp,
	}
}

func (skiplist *SkipList) RandomHeight() int {
	height := 1
	for {
		rnd, _ := rand.Int(rand.Reader, big.NewInt(64))
		if height < kBranching && rnd.Int64()%kBranching == 0 {
			height++
		} else {
			break
		}
	}
	return height
}

func (skiplist *SkipList) KeyIsAfterNode(key interface{}, n *Node) bool {
	return (n != nil) && (skiplist.compare.Compare(n.key, key) < 0)
}

func (skiplist *SkipList) GetMaxHeight() int {
	return int(atomic.LoadInt64(&skiplist.maxHeight))
}

func (skiplist *SkipList) Equal(a, b interface{}) bool {
	return skiplist.compare.Compare(a, b) == 0
}

func (skiplist *SkipList) FindGreaterOrEqual(key interface{}) ([]*Node, *Node) {
	x := skiplist.head
	level := skiplist.GetMaxHeight() - 1
	prev := make([]*Node, kMaxHeight)
	for {
		next := x.Next(level)
		if skiplist.KeyIsAfterNode(key, next) {
			x = next
		} else {
			prev[level] = x
			if level == 0 {
				return prev, next
			} else {
				// Switch to next list
				level--
			}
		}
	}
}

func (skiplist *SkipList) FindLessThan(key interface{}) *Node {
	x := skiplist.head
	level := skiplist.GetMaxHeight() - 1
	for {
		next := x.Next(level)
		if next == nil || skiplist.compare.Compare(next.key, key) >= 0 {
			if level == 0 {
				return x
			} else {
				// Switch to next list
				level--
			}
		} else {
			x = next
		}
	}
}

func (skiplist *SkipList) FindLast() *Node {
	x := skiplist.head
	level := skiplist.GetMaxHeight() - 1
	for {
		next := x.Next(level)
		if next == nil {
			if level == 0 {
				return x
			} else {
				level--
			}
		} else {
			x = next
		}
	}
}

func (skiplist *SkipList) Insert(key interface{}) {
	prev, x := skiplist.FindGreaterOrEqual(key)
	height := skiplist.RandomHeight()
	if height > skiplist.GetMaxHeight() {
		for i := skiplist.GetMaxHeight(); i < height; i++ {
			prev[i] = skiplist.head
		}
		atomic.StoreInt64(&skiplist.maxHeight, int64(height))
	}
	x = newNode(key, height)
	for i := 0; i < height; i++ {
		x.SetNext(i, prev[i].Next(i))
		prev[i].SetNext(i, x)
	}
}

func (skiplist *SkipList) Contains(key interface{}) bool {
	_, x := skiplist.FindGreaterOrEqual(key)
	return x != nil && skiplist.Equal(key, x.key)
}

func (skiplist *SkipList) EstimateCount(key interface{}) int {
	count := 0
	x := skiplist.head
	level := skiplist.GetMaxHeight() - 1
	for {
		next := x.Next(level)
		if next == nil || skiplist.compare.Compare(next.key, key) >= 0 {
			if level == 0 {
				return count
			} else {
				count = count * kBranching
				level--
			}
		} else {
			x = next
			count++
		}
	}
}
