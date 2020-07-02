package easy

/*
leetcode: 203
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

//RemoveElements ...
func RemoveElements(head *ListNode, val int) *ListNode {
	cur := head

	back := &ListNode{Next: cur}

	for cur != nil {
		if cur.Val == val {
			if cur == head {
				head = head.Next
				cur = head
				back.Next = cur
			} else {
				back.Next = cur.Next
				cur = cur.Next
			}
		} else {
			back = cur
			cur = cur.Next
			back.Next = cur
		}
	}
	return head
}
