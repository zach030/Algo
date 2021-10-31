package concurrent

import (
	"fmt"
	"sync"
)

// 两个线程交替打印1-100
func concurrentPrint() {
	c1, c2,c3 := make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)
	c1 <- struct{}{}
	count := 1
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		for count <= 98 {
			select {
			case <-c1:
				fmt.Println("t-1: ", count)
				count++
				c2 <- struct{}{}
			}
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		for count <= 99 {
			select {
			case <-c2:
				fmt.Println("t-2: ", count)
				count++
				c3 <- struct{}{}
			}
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		for count <= 100 {
			select {
			case <-c3:
				fmt.Println("t-3: ", count)
				count++
				c1 <- struct{}{}
			}
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("done")
}
