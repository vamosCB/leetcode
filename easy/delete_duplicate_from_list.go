package easy

import "github.com/vamosCB/utils/list"

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

*/

//DeleteDuplicates ...
func DeleteDuplicates(head *list.Node) *list.Node {

	output := head

	pre := head

	for head != nil {
		if head.Data != pre.Data {
			pre.Next = head
			pre = head
		}
		head = head.Next
	}
	if pre != head {
		pre.Next = head
	}
	return output
}
