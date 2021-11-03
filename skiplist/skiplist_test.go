package skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList_Search(t *testing.T) {
	sk := NewSkipList()
	initRelation()
	entry := sk.Search([]byte("3"))
	fmt.Printf("key: %v, value: %v", string(entry.Key), string(entry.Value))
}
