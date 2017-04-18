/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/13 14:57
  */

package main

import (
	"Golang_Algorithm/Data_Structure/Queue"
	"fmt"
	"log"
)

func main() {
	testQueue := Queue.New()
	testQueue.PrintQueue()
	fmt.Println("Length is", testQueue.LenQueue())
	fmt.Println("isEmpty:", testQueue.IsEmpty())

	for i := 1; i <= 10; i++ {
		testQueue.Enqueue(i)
	}
	testQueue.PrintQueue()
	fmt.Println("Length is", testQueue.LenQueue())
	fmt.Println("isEmpty:", testQueue.IsEmpty())

	ret, err := testQueue.Pos(3)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(ret)
	}

	for i := 0; i < 15; i++ {
		item, err := testQueue.Dequeue()
		if err != nil {
			log.Println(err)
			break
		}
		if item != i + 1 {
			fmt.Println("Error")
			break
		}
		fmt.Println("Dequeue", item)
		if !testQueue.IsEmpty() {
			fmt.Println("Length is", testQueue.LenQueue())
		} else {
			fmt.Println("Queue is Empty")
		}
	}
}
