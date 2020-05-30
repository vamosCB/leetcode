package main

import (
	"leetcode/own"
	"utils/list"
)

func main() {

	node1 := &list.Node{Data: 1}
	node2 := &list.Node{Data: 2}
	node3 := &list.Node{Data: 3}

	node1.Next = node2
	node2.Next = node3
	testList := &list.SingleLinkedList{Head: node1}

	own.PrintListReverse2(testList.Head)
}
