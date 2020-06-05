package main

import (
	"fmt"
	"leetcode/easy"
	"utils/list"
)

func main() {

	nums1 := []int{0, 0, 0, 0, 0, 0}
	nums2 := []int{1}
	easy.MergeTwoSortArray(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}

func returnTestList() *list.Node {
	node9 := &list.Node{Data: 3}
	node8 := &list.Node{Data: 3, Next: node9}
	node7 := &list.Node{Data: 3, Next: node8}
	node6 := &list.Node{Data: 3, Next: node7}
	node5 := &list.Node{Data: 3, Next: node6}
	node4 := &list.Node{Data: 2, Next: node5}
	node3 := &list.Node{Data: 2, Next: node4}
	node2 := &list.Node{Data: 1, Next: node3}
	node1 := &list.Node{Data: 1, Next: node2}
	return node1
}
