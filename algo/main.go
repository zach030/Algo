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

func main() {
	var group int
	fmt.Scan(&group)
	groups := make([]taskGroup, 0, group)
	for i := 0; i < group; i++ {
		var num int
		fmt.Scan(&num)
		tasks := make(map[int]*task, num)
		noDep := make(map[int]*task, 0)
		for j := 1; j <= num; j++ {
			t := &task{id: j}
			tasks[t.id] = t
			fmt.Scan(&t.day, &t.dNum)
			t.depends = make(map[int]*task)
			t.children = make(map[int]*task)
			if t.dNum == 0 {
				noDep[t.id] = t
			} else {
				for k := 0; k < t.dNum; k++ {
					var dID int
					fmt.Scan(&dID)
					t.depends[dID] = tasks[dID]
					tasks[dID].children[t.id] = t
				}
			}
		}
		groups = append(groups, taskGroup{
			num:     num,
			allTask: tasks,
			noDep:   noDep,
		})
	}

	startPlan(groups)
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


func TestSolution3(t *testing.T){
	var n,m int
	fmt.Scan(&n,&m)
	var height int
	fmt.Scan(&height)
	var query int
	fmt.Scan(&query)

}