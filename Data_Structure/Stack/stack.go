/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/13 15:27
  */

package Stack

import (
	"sync"
	"fmt"
	"errors"
)

type stackNode struct {				// 栈节点结构
	stackEle 	interface{}
	next 		*stackNode
}

type Stack struct {					// 栈结构
	top		 	*stackNode
	depth 		int
	lock 		*sync.Mutex
}

// 创建一个栈
func New() *Stack {
	newStack := new(Stack)
	newStack.depth = 0
	newStack.lock = new(sync.Mutex)
	return newStack
}

// 元素入栈
func (stack *Stack) Push(element interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	{
		newEle := new(stackNode)
		newEle.stackEle = element
		newEle.next = stack.top
		stack.top = newEle
		stack.depth ++
	}
}

// 元素出栈
func (stack *Stack) Pop() (interface{}, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	{
		if stack.depth > 0 {
			element := stack.top.stackEle
			stack.top = stack.top.next
			stack.depth --
			return element, nil
		}
		return nil, errors.New("Error: Stack is Empty")
	}
}

// 获取栈深度
func (stack *Stack) DepthStack() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.depth
}

// 判断栈是否为空
func (stack *Stack) IsEmpty() bool {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.depth == 0
}

// 自顶向下遍历栈
func (stack *Stack) PrintStack()  {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	{
		if stack.depth == 0 {
			fmt.Println("Stack is Empty")
			return
		}
		var temp = stack.top
		for i := 0; i < stack.depth; i++ {
			fmt.Printf("%v ", temp.stackEle)
			temp = temp.next
		}
		fmt.Println()
	}
}

// 返回栈顶元素
func (stack *Stack) Top() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	{
		return stack.top.stackEle
	}
}