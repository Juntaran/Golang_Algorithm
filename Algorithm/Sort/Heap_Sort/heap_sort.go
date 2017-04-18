/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 17:19
  */

package Heap_Sort

import "Golang_Algorithm/Data_Structure/Heap"

func Heap_Sort(nums []int)  {
	sortHeap := heap.NewMin()
	for i := 0; i < len(nums); i++ {
		sortHeap.Insert(heap.Int(nums[i]))
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = int(sortHeap.Extract().(heap.Int))
	}
}