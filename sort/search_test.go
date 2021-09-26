package sort

import (
	"fmt"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	arr := []int{1, 4, 6, 7, 9, 12}
	target := 6
	fmt.Println(binarySearch(arr, target))

	fmt.Println(bsSearchInternal(arr, 0, len(arr)-1, target))
}

func TestFindKMax(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	findKMax(nums, 0, len(nums)-1, 4)
	findKMax2(nums, 0, len(nums)-1, 4)

	arr := []int{5, 3, 5, 4}
	fmt.Println(partition(arr, 0, len(arr)-1))
}

func findKthLargest(nums []int, k int) int {
	return findKMax2(nums,0,len(nums)-1,k)
}

func findKMax2(nums []int,left,right,k int)int{
	pos := partition2(nums,left,right)
	if pos == k{
		return nums[pos]
	}else if pos>k{
		return partition2(nums,pos+1,right)
	}
	return partition2(nums,left,pos-1)
}

func partition2(nums []int,left,right int)int{
	tmp := nums[left]
	for left<right{
		for left<right && nums[left]<=tmp{
			left++
		}
		nums[left]=nums[right]
		for left<right && nums[right]>=tmp{
			right--
		}
		nums[right]=nums[left]
	}
	nums[left]=tmp
	return right
}
