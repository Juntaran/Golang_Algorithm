/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/15 10:51
  */

package main

import (
	"Golang_Algorithm/Data_Structure/Heap"
	"fmt"
	"log"
)

func main() {
	// 最小堆测试
	h1 := heap.NewMin()
	h1.Insert(heap.Int(8))
	h1.Insert(heap.Int(7))
	h1.Insert(heap.Int(6))
	h1.Insert(heap.Int(3))
	h1.Insert(heap.Int(1))
	h1.Insert(heap.Int(0))
	h1.Insert(heap.Int(2))
	h1.Insert(heap.Int(4))
	h1.Insert(heap.Int(9))
	h1.Insert(heap.Int(5))

	sorted := make([]heap.Int, 0)
	for h1.Size() > 0 {
		sorted = append(sorted, h1.Extract().(heap.Int))
	}
	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i] > sorted[i+1] {
			fmt.Println(sorted)
		} else {
			log.Println("Error")
		}
	}

	h2 := heap.NewMin()
	h2.Insert(heap.Int(8))
	h2.Insert(heap.Int(7))
	h2.Insert(heap.Int(6))
	h2.Insert(heap.Int(3))
	h2.Insert(heap.Int(1))
	h2.Insert(heap.Int(0))
	h2.Insert(heap.Int(2))
	h2.Insert(heap.Int(4))
	h2.Insert(heap.Int(9))
	h2.Insert(heap.Int(5))

	sorted = make([]heap.Int, 0)
	for h2.Size() > 0 {
		sorted = append(sorted, h2.Extract().(heap.Int))
	}
	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i] > sorted[i+1] {
			fmt.Println(sorted)
		} else {
			log.Println("Error")
		}
	}
}