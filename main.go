package main

import (
	"fmt"

	"github.com/vamosCB/utils/list"
)

func main() {
	aaa := []int{1, 2, 3}
	fmt.Println(append([]int{4, 5, 6}, aaa...))
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
