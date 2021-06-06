package Algo

import (
	"fmt"
	"testing"
)


func TestLRU_Add(t *testing.T) {
	lru := NewLRU(3)
	lru.Add("1",ele{
		key:   "1",
		value: "hello",
	})
	lru.Add("2",ele{
		key:   "2",
		value: "world",
	})
	lru.Add("3",ele{
		key:   "3",
		value: "there",
	})
	fmt.Println(lru.Get("1"))
}
