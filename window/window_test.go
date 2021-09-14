package window

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func TestFindUnsortedSubarray(t *testing.T) {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{2, 1}))
	fmt.Println(findUnsortedSubarray([]int{5, 4, 3, 2, 1}))
}

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func TestFindArgs(t *testing.T) {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("baa", "aa"))
}

func TestCalcNum(t *testing.T) {
	fmt.Println(calcNum(1, 1, 1))
	fmt.Println(calcNum(1, 1, 2))
	fmt.Println(calcNum(3, 4, 5))
}

func calcNum(x, y, r int) int {
	if x == y {
		if r < x {
			return 0
		} else if r == x {
			return 2
		} else if r > x {
			return 4
		}
	}
	min := min(x, y)
	max := max(x, y)
	if r*r == x*x+y*y {
		return 3
	}
	if r < min {
		return 0
	} else if r == min {
		return 1
	} else if r > min && r < max {
		return 2
	} else if r == max {
		return 3
	} else {
		return 4
	}

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestChooseSweet(t *testing.T) {
	arr := []int{1, 1, 1, 1}
	max := 2
	num := 2
	fmt.Println(chooseSweet(arr, max, num))
}

// 从arr里选出连续num个小于等于max的
func chooseSweet(arr []int, max, num int) int {
	_, right := 0, 0
	ret := 0
	window := make(map[int]int)
	for right < len(arr) {
		if arr[right] <= max {
			window[arr[right]]++
		}
		right++
		if len(window) == num {
			ret++
			window = make(map[int]int)
		}
	}
	return ret
}

func TestFindSubStr(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring2("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring2("bbbb"))
	fmt.Println(lengthOfLongestSubstring2("pwwkew"))
}

func TestSolution1(t *testing.T) {
	fmt.Println(string(findKthBit(4, 11)))
}

func calc(arr []int64, target int64) int {
	total := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] < target {
				total++
			}
		}
	}
	return total
}

var (
	s1     = "a"
	s2     = "abz"
	s3     = "abzcayz"
	s4     = "abzcayzdabzxayz"
	str    = []string{s1, s2, s3, s4}
	letter = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

//
func findKthBit(n int, k int) byte {
	str := GetSn(n)
	return str[k-1]
}

func GetSn(n int) string {
	if n == 1 {
		return s1
	}
	if n == 2 {
		return s2
	}
	prev := GetSn(n - 1)
	alpha := letter[n-1]
	afterInvert := make([]byte, 0, len(prev))
	for i := 0; i < len(prev); i++ {
		idx := 25 - (prev[i] - 97)
		afterInvert = append(afterInvert, letter[idx]...)
	}
	afterReverse := make([]byte, 0, len(afterInvert))
	for i := len(afterInvert) - 1; i >= 0; i-- {
		afterReverse = append(afterReverse, afterInvert[i])
	}
	sum := prev + alpha + string(afterReverse)
	fmt.Println("s", n, "is:", sum)
	return sum
}

func TestSolution2(t *testing.T) {
	arr := []int{4, 4, 5}
	arr2 := []int{1, 1, 1}
	arr3 := []int{1, 2, 3}
	arr4 := []int{5}

	fmt.Println(GreddyDispatch(arr))
	fmt.Println(GreddyDispatch(arr2))
	fmt.Println(GreddyDispatch(arr3))
	fmt.Println(GreddyDispatch(arr4))
}

func GreddyDispatch(age []int) int {
	dis := make([]int, len(age))
	for i := 0; i < len(age); i++ {
		dis[i] = 1
	}
	for i := 1; i < len(age); i++ {
		if age[i] > age[i-1] {
			dis[i] = dis[i-1] + 1
		}
	}
	for i := len(age) - 1; i > 0; i-- {
		if age[i] < age[i-1] {
			dis[i-1] = max(dis[i-1], dis[i]+1)
		}
	}
	sum := 0
	for _, di := range dis {
		sum += di
	}
	return sum
}

func TestSolution3(t *testing.T) {
	fmt.Println(minSailCost([][]int{{1, 1, 1, 1, 0}, {0, 1, 0, 1, 0}, {1, 1, 2, 1, 1}, {0, 2, 0, 0, 1}}))
}

func minSailCost(input [][]int) int {
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		down := dfs(x+1, y)
		right := dfs(x, y+1)
		if x >= len(input) || y >= len(input[x]) || input[x][y] == 2 {
			return -1
		}
		if input[x][y] == 1 {
			return 1
		}
		if input[x][y] == 0 {
			return 2
		}
		if down < right {
			return down
		}
		return right
	}
	return dfs(0, 0)
}

var input = []int{1, 1, 2}
var limit = make(map[string]bool, 0)

func TestMeituanSolution1(t *testing.T) {

	arr := []int{1, 2, 3}
	fullSort(arr)
}

func search(arr []int) int {
	size := len(arr)
	front := size - 2
	for front >= 0 && arr[front] > arr[front+1] {
		front--
	}
	return front
}

func searchBigger(arr []int, targetIndex int) int {
	size := len(arr)
	index := size - 1
	for index >= 0 && arr[index] < arr[targetIndex] {
		index--
	}
	return index
}

func swap(arr []int, id1, id2 int) {
	arr[id1], arr[id2] = arr[id2], arr[id1]
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func printArray(arr []int) {
	str := ""
	for i := 0; i < len(arr); i++ {
		if str == "" {
			str = fmt.Sprintf("%s%v", str, input[arr[i]-1])
		} else {
			str = fmt.Sprintf("%s %v", str, input[arr[i]-1])
		}
	}
	if _, ok := limit[str]; ok {
		return
	}
	limit[str] = true
	fmt.Println(str)
}

func fullSort(arr []int) {
	if arr == nil || len(arr) == 0 {
		return
	}
	printArray(arr)
	front := search(arr)
	for front != -1 {
		back := searchBigger(arr, front)
		swap(arr, front, back)
		reverse(arr, front+1, len(arr)-1)
		printArray(arr)
		front = search(arr)
	}
}

//1 a
//1 b
//2 6
//2 4

func TestMTSolution2(t *testing.T) {
	type op struct {
		typ int
		arg interface{}
	}
	str := "asdgfas"
	opArr := []op{
		{
			typ: 2,
			arg: 6,
		},
		{
			typ: 1,
			arg: "a",
		},
		{
			typ: 1,
			arg: "b",
		},
		{
			typ: 2,
			arg: 6,
		},
		{
			typ: 2,
			arg: 4,
		},
	}
	for _, op := range opArr {
		switch op.typ {
		case 1:
			str = fmt.Sprintf("%s%v", str, op.arg.(string))
		case 2:
			fmt.Println(query(str, op.arg.(int)-1))
		}
	}
}

func query(str string, pos int) int {
	size := len(str)
	step := size
	found := false
	for i := pos + 1; i < size; i++ {
		if str[i] == str[pos] {
			found = true
			if i-pos < step {
				step = i - pos
			}
		}
	}
	for i := pos - 1; i >= 0; i-- {
		if str[i] == str[pos] {
			found = true
			if pos-i < step {
				step = pos - i
			}
		}
	}
	if found {
		return step
	}
	return -1
}

func TestBaiduSolution1(t *testing.T) {
	_, k := 2, 2
	matrix := [][]int{{0, 1}, {1, 0}}
	fmt.Println(bigger(k, matrix))
}

func bigger(k int, matrix [][]int) [][]int {
	e := len(matrix)
	bigM := make([][]int, k*e)
	for i := 0; i < k*e; i++ {
		bigM[i] = make([]int, k*e)
	}
	for i := 0; i < len(bigM); i++ {
		for j := 0; j < len(bigM); j++ {
			hi,hj:=i/k,j/k
			bigM[i][j]=matrix[hi][hj]
		}
	}
	return bigM
}

func TestBaiduSolution2(t *testing.T) {
	n := 30
	eulerInit()
	fmt.Println(phi[n])
}

var phi = make([]int, 1010)

func eulerInit() {
	top := 1000
	for i := 2; i <= top; i++ {
		phi[i] = i
	}
	for i := 2; i <= top; i++ {
		if phi[i] == i {
			for j := i; j <= top; j += i {
				phi[j] -= phi[j] / i
			}
		}
		phi[i] += phi[i-1]
	}
	fmt.Println(phi)
}

// n以内互为质数的对数
func getPair(n int64) int64 {
	res := n
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			res = (res / i) * (i - 1)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n != 1 {
		res = (res / n) * (n - 1)
	}
	return res
}
