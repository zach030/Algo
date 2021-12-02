package daily

type ListNode struct {
	Val  int
	Next *ListNode
}

//class Solution {
//    ArrayList<Integer> tmp = new ArrayList<Integer>();
//    public int[] reversePrint(ListNode head) {
//        recur(head);
//        int[] res = new int[tmp.size()];
//        for(int i = 0; i < res.length; i++)
//            res[i] = tmp.get(i);
//        return res;
//    }
//    void recur(ListNode head) {
//        if(head == null) return;
//        recur(head.next);
//        tmp.add(head.val);
//    }
//}
//
func reversePrint(head *ListNode) []int {
	tmp := make([]int, 0)
	var recur = func(node *ListNode) {}
	recur = func(node *ListNode) {
		if node == nil {
			return
		}
		recur(node.Next)
		tmp = append(tmp, node.Val)
	}
	recur(head)
	return tmp
}
