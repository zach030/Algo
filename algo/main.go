package main

import (
	"fmt"
	"testing"
)

type taskGroup struct {
	num     int
	noDep   map[int]*task
	allTask map[int]*task
}

type task struct {
	id       int
	day      int
	dNum     int
	depends  map[int]*task
	children map[int]*task
}

var match = map[int]int{0: 1, 2: 3}

func main() {
	nums := []int{0, 0, 0, 1, 1, 1}
	count := 0
	stack := make([]int, 0)
	stack = append(stack, nums[0])
	for len(stack) != 0 {
		for i := 1; i < len(nums); i++ {
			if nums[i] == 0 || nums[i] == 2 {
				stack = append(stack, nums[i])
				continue
			}
			if nums[i] == match[stack[len(stack)-1]] {
				count += 2
				stack = stack[:len(stack)-1]
				continue
			}
		}
	}
	fmt.Println(count)
}

func dis(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func calcCost(a, b []int, x, y, z int) int {
	a1, a2 := a[0], a[1]
	b1, b2 := b[0], b[1]
	if a1 == b1 {
		return x*dis(a2, b2) + z
	}
	if a2 == b2 {
		return x*dis(a1, b1) + z
	}
	return x*(dis(a1, b1)+dis(a2, b2)) + y + z
}

func startPlan(groups []taskGroup) {
	for i := 0; i < len(groups); i++ {
		g := groups[i]
		startEachPlan(g)
	}
}

func startEachPlan(g taskGroup) {
	days := 0
	for id, noDepTask := range g.noDep {
		if noDepTask.day > days {
			days = noDepTask.day
		}
		for _, child := range noDepTask.children {
			delete(child.depends, id)
		}
		delete(g.allTask, id)
	}
	for len(g.allTask) != 0 {
		for id, task := range g.allTask {
			if len(task.depends) == 0 {
				if task.day > days {
					days += task.day
				}
				for _, child := range task.children {
					delete(child.depends, id)
				}
			}
			delete(g.allTask, id)
		}
	}
	fmt.Println(days)
}

func getMinest(a, b, c int64) int64 {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if a >= b {
		if b < c {
			return b
		} else {
			return c
		}
	}
	return 0
}

func getMin(arr []int64) ([]int64, int64) {
	ele := arr[0]
	arr = arr[1:]
	return arr, ele
}

func SelectSort(arr []int64) []int64 {
	for i := 0; i < len(arr); i++ {
		pos := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[pos] {
				pos = j
			}
		}
		arr[pos], arr[i] = arr[i], arr[pos]
	}
	return arr
}

func TestSolution3(t *testing.T) {
	var n, m int
	fmt.Scan(&n, &m)
	var height int
	fmt.Scan(&height)
	var query int
	fmt.Scan(&query)

}

func f(p map[string][]string) {
	for i, c := range topoSort(p) {
		fmt.Printf("%d:\t%s\n", i+1, c)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	visited := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	//sort.Strings(keys)
	visitAll(keys)
	return order
}
