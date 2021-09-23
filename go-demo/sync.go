package go_demo

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
	age  int
}

var personPool = sync.Pool{
	New: func() interface{} {
		return Person{}
	},
}

func syncDemo() {
	p := personPool.Get()
	person := p.(Person)
	person.name = "zach"
	person.age = 21
	personPool.Put(p)

	p1 := personPool.Get()
	fmt.Printf("%+v", p1.(Person))
}

