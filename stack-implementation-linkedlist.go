// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
)

type ListNode struct {
	data interface{}
	next *ListNode
}

type Stack struct {
	top  *ListNode
	size int
}

func (stk *Stack) IsFull() bool {
	return false
}

func (stk *Stack) IsEmpty() bool {
	return stk.size == 0
}

func (stk *Stack) Push(val int) {
	stk.top = &ListNode{data: val, next: stk.top}
	stk.size++
}

func (stk *Stack) Pop() error {
	if stk.IsEmpty() {
		return errors.New("Empty Stack")
	}
	stk.size--
	stk.top = stk.top.next
	return nil
}

func (stk *Stack) Peek() (interface{}, error) {
	if stk.IsEmpty() {
		return nil, errors.New("Empty Stack")
	}
	return stk.top.data, nil
}

func main() {
	stk := new(Stack)

	for i := 1; i <= 5; i++ {
		stk.Push(i)
	}

	stk.Pop()
	data, err := stk.Peek()

	if err != nil {
		fmt.Printf("Peek failed with error:%v", err)
		fmt.Println()
		return
	}

	fmt.Printf("Top element is : %v", data)
	fmt.Println()
	
	
       // Traversing Stack
	p := stk.top
	for i := 1; i <= stk.size; i++ {
		fmt.Println(p.data)
		p = p.next

	}
}
