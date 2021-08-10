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