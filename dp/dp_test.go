package dp

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("babad"))
}

func TestFib(t *testing.T) {
	fmt.Println(fib(95))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestReverseList(t *testing.T) {
	l := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
		},
	}
	k := 2
	i := 0
	for i < k {
		reverseN(l, k*i, (1+i)*k)
		i++
	}
	n := l
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

func reverseN(node *ListNode, l, r int) *ListNode {
	head := node
	left := node
	right := node
	i := 0
	for i < r {
		if i == l {
			left = head
		}
		head = head.Next
		i++
	}
	right = head
	var reverse func(node, tail *ListNode) *ListNode
	reverse = func(head, tail *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		var prev *ListNode
		for head.Next != tail {
			tmp := head.Next
			head.Next = prev
			prev = head
			head = tmp
		}
		return prev
	}
	return reverse(left, right)
}

func TestLongestLIS(t *testing.T) {
	a := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	fmt.Println(a)
}

func TestMinStep(t *testing.T) {
	s := minSteps(3)
	fmt.Println(s)
}

func TestMinCost(t *testing.T) {
	a := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	fmt.Println(a)

	b := uniquePaths(3, 3)
	fmt.Println(b)
}

func TestDeleteAndRun(t *testing.T) {
	fmt.Println(deleteAndEarn([]int{2, 2, 4, 3, 3, 6, 7}))

	//fmt.Println(canPartition([]int{1, 5, 5, 11}))

	camelCase("")
}

func camelCase(input string) {
	fmt.Println(isMatch("a+b", "a*b"))
}

func isMatch(s, p string) bool {
	n, m := len(s), len(p)
	f := make([][]bool, n+1)

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if j == 0 {
				f[i][j] = i == 0
			} else {
				if p[j-1] != '*' {
					if i > 0 && s[i-1] == p[j-1] || p[j-1] == '.' {
						f[i][j] = f[i-1][j-1]
					}
				} else {
					if j >= 2 {
						res := f[i][j] || f[i][j-2]
						f[i][j] = res
					}
					if i >= 1 && j >= 2 && s[i-1] == p[j-2] || p[j-2] == '.' {
						res := f[i][j] || f[i-1][j]
						f[i][j] = res
					}
				}
			}
		}
	}
	return f[n][m]
}

type round struct {
	off      int
	def      int
	monsters []monster
}

type monster struct {
	off  int
	def  int
	life int
}

func TestWar1(t *testing.T) {
	r := []round{
		{
			def: 1,
			off: 1,
			monsters: []monster{
				{off: 10, def: 5, life: 5},
			},
		},
		{
			off: 10,
			def: 2,
			monsters: []monster{
				{off: 12, def: 6, life: 9},
				{off: 8, def: 9, life: 8},
				{off: 11, def: 5, life: 4},
			},
		},
	}
	startWar(r)
}

func startWar(rounds []round) {
	for _, round := range rounds {
		failed := false
		lifeSpan := make([]int, 0)
		for _, monster := range round.monsters {
			// 不可能通过
			if round.off <= monster.def {
				failed = true
				break
			}
			// 对某一只怪兽
			for monster.life > 0 {
				mHurt := round.off - monster.def
				monster.life -= mHurt
				if monster.life <= 0 {
					break
				}
				iHurt := monster.off - round.def
				if iHurt <= 0 {
					iHurt = 0
				}
				lifeSpan = append(lifeSpan, 0-iHurt)
			}
			lifeSpan = append(lifeSpan, 0-monster.life)
		}
		if failed {
			fmt.Println(-1)
		} else {
			fmt.Println(lifeSpan)
			min := lifeSpan[0]
			subMin := 0
			for _, i2 := range lifeSpan {
				if subMin <= 0 {
					subMin += i2
				} else {
					subMin = i2
				}
				if subMin < min {
					min = subMin
				}
			}
			fmt.Println(0 - min + 1)
		}
	}
}

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

func NetSolution2() {
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

	fmt.Println(groups)
}

func TestSolution2(t *testing.T) {
	task1_1 := &task{id: 1, day: 5, dNum: 0}
	task1_2 := &task{id: 2, day: 10, dNum: 1, depends: map[int]*task{1: task1_1}}
	task1_3 := &task{id: 3, day: 15, dNum: 1, depends: map[int]*task{1: task1_1}}

	task2_1 := &task{id: 1, day: 3, dNum: 0}
	task2_2 := &task{id: 2, day: 4, dNum: 0}
	task2_3 := &task{id: 3, day: 7, dNum: 1, depends: map[int]*task{1: task2_1}}
	task2_4 := &task{id: 4, day: 6, dNum: 2, depends: map[int]*task{1: task2_1, 2: task2_1}}
	groups := []taskGroup{
		{
			num: 3,
			allTask: map[int]*task{
				1: task1_1,
				2: task1_2,
				3: task1_3,
			},
		},
		{
			num: 4,
			allTask: map[int]*task{
				1: task2_1,
				2: task2_2,
				3: task2_3,
				4: task2_4,
			},
		},
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
	remain := len(g.allTask)
	for id, noDepTask := range g.noDep {
		if noDepTask.day > days {
			days = noDepTask.day
		}
		for _, child := range noDepTask.children {
			delete(child.depends, id)
		}
		delete(g.allTask, id)
	}
	for remain != 0 {
		for id, task := range g.allTask {
			if len(task.depends) == 0 {
				if task.day > days {
					days = task.day
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

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))
}

func TestSocks(t *testing.T) {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
}

func TestBestSight(t *testing.T) {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
	fmt.Println(maxScoreSightseeingPair([]int{1, 2}))
	fmt.Println(maxScoreSightseeingPair([]int{7, 8, 8, 10}))
	fmt.Println(maxScoreSightseeingPair([]int{7, 2, 6, 6, 9, 4, 3}))
}

func TestMinPath(t *testing.T) {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func TestFindNumInMatrix(t *testing.T) {
	findNumberIn2DArray([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20)
}

func TestReplaceString(t *testing.T) {
	fmt.Println(replaceSpace("how are you"))
}
