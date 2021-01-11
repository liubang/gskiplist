package gskiplist

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	skiplist := NewSkipList(NewIntComparator())
	for i := 0; i < 1000; i++ {
		skiplist.Insert(rand.Int())
	}
	it := NewIterator(skiplist)
	i := 0
	for it.SeekToFirst(); it.Valid(); it.Next() {
		i++
		fmt.Println(it.Key(), "::", skiplist.EstimateCount(it.Key()))
	}
	fmt.Println(i)
}
