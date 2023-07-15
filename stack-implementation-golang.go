// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top      int
	capacity int
	data     []interface{}
}

func (stk *Stack) Init(capacity int) *Stack {
	stk.top = -1
	stk.capacity = capacity
	stk.data = make([]interface{}, capacity)
	return stk
}

func NewStack(capacity int) *Stack {
	stk := Stack{}
	return stk.Init(capacity)
}

func (stk *Stack) StackIsFull() bool {
	if stk.top == stk.capacity-1 {
		return true
	}
	return false
}

func (stk *Stack) Push(val interface{}) error {
	if stk.StackIsFull() {
		return errors.New("Stack is full")
	}
	stk.top++
	stk.data[stk.top] = val
	return nil
}

func (stk *Stack) StackIsEmpty() bool {
	return stk.top == -1
}

func (stk *Stack) Pop() error {
	if stk.StackIsEmpty() {
		return errors.New("Stack is Empty")
	}
	stk.data[stk.top] = nil
	stk.top--
	return nil
}

func (stk *Stack) Peek() (interface{}, error) {
	if stk.StackIsEmpty() {
		return nil, errors.New("Stack is Empty")
	}
	temp := stk.data[stk.top]
	return temp, nil
}

func main() {
	stk := NewStack(10)
	for i := 1; i <= 5; i++ {
		stk.Push(i)
	}

	stk.Pop()
	temp, err := stk.Peek()
	if err != nil {
		fmt.Printf("Peek failed with error : %v", err)
		return
	}
	fmt.Printf("Current element at top is: %v", temp)
}
