package daily

import (
	"fmt"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	list := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
	}
	fmt.Println(reversePrint(list))
}
