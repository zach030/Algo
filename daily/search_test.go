package daily

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_missingNumber(t *testing.T) {
	arr1 := []int{0, 1, 3}
	fmt.Println(missingNumber(arr1))
	arr2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(missingNumber(arr2))
}

func TestFindInMatrix(t *testing.T) {
	matrix := [][]int{{1, 2}, {3, 4}}
	fmt.Println(matrix[1][0])
	fmt.Println(findNumberIn2DArray(matrix, 2))
}

func TestMinArr(t *testing.T) {
	a := minArray([]int{3, 4, 5, 1, 2})
	fmt.Println(a)
}

func TestUniqCha(t *testing.T) {
	fmt.Println(string(firstUniqChar("loveleetcode")))
}

func TestLevelOrder(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(levelOrder(tree))
}

func TestSearchWords(t *testing.T) {
	m := [][]byte{{byte('A'), byte('B'), byte('C'), byte('E')},
		{byte('S'), byte('F'), byte('C'), byte('S')},
		{byte('A'), byte('D'), byte('E'), byte('E')}}
	fmt.Println(exist(m, "ABCCEE"))
}

func Test_movingCount(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				m: 2,
				n: 3,
				k: 1,
			},
			want: 3,
		},
		{
			name: "t2",
			args: args{
				m: 3,
				n: 1,
				k: 0,
			},
			want: 1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
