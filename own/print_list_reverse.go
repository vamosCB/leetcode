package own

import (
	"fmt"

	"github.com/vamosCB/utils/list"
	"github.com/vamosCB/utils/stack"
)

/*
	倒叙打印单链表的值

	  1. 不能修改原本链表结构(不能反转链表)
	  2. 对于逆序想到的第一思路就是用栈（先进后出）
*/

//PrintListReverse  正常思路
func PrintListReverse(l1 *list.Node) {
	if l1 == nil {
		return
	}

	stack := stack.InitStack()
	head := l1
	for head != nil {
		stack.Push(head.Data)
		head = head.Next
	}

	for !stack.IsEmpty() {
		fmt.Print(stack.Top())
		stack.Pop()
	}
}

//PrintListReverse2 使用递归
func PrintListReverse2(l1 *list.Node) {
	head := l1
	if head != nil {
		if head.Next != nil {
			PrintListReverse2(head.Next)
		}
		fmt.Print(head.Data)
	}
}
