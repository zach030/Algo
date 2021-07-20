package window

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func TestFindUnsortedSubarray(t *testing.T) {
	fmt.Println(findUnsortedSubarray([]int{2,6,4,8,10,9,15}))
	fmt.Println(findUnsortedSubarray([]int{1,2,3,4}))
	fmt.Println(findUnsortedSubarray([]int{2,1}))
	fmt.Println(findUnsortedSubarray([]int{5,4,3,2,1}))
}