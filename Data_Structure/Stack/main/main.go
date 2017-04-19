/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/13 15:52
  */

package main

import (
	"Golang_Algorithm/Data_Structure/Stack"
	"fmt"
	"log"
)

func main() {
	testStack := Stack.New()
	testStack.PrintStack()
	fmt.Println("Depth is", testStack.DepthStack())
	fmt.Println("isEmpty", testStack.IsEmpty())

	for i := 1; i <= 10; i++ {
		testStack.Push(i)
	}
	testStack.PrintStack()
	fmt.Println("Length is", testStack.DepthStack())
	fmt.Println("isEmpty", testStack.IsEmpty())

	fmt.Println(testStack.Top())

	for i := 0; i < 15; i++ {
		item, err := testStack.Pop()
		if err != nil {
			log.Println(err)
			break
		}
		if item != 10 - i {
			fmt.Println("Error")
			break
		}
		fmt.Println("Pop", item)
		if !testStack.IsEmpty() {
			fmt.Println("Length is", testStack.DepthStack())
		} else {
			fmt.Println("Stack is Empty")
		}
	}
}
