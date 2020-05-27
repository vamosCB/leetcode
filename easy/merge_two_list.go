package easy

/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

//ListNode node
type ListNode struct {
	Val  int
	Next *ListNode
}

//MergeTwoLists 合并两个有序链表为一个有序链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList = &ListNode{}
	var out = newList
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			newList.Next = l1
			l1 = l1.Next
			newList = newList.Next
		} else {
			newList.Next = l2
			l2 = l2.Next
			newList = newList.Next
		}
	}
	if l1 != nil {
		newList.Next = l1
	} else if l2 != nil {
		newList.Next = l2
	}

	return out.Next
}
