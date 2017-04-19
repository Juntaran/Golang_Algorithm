/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 17:19
  */

package Heap_Sort

//import "Golang_Algorithm/Data_Structure/Heap"

func Heap_Sort(nums []int)  {
	//sortHeap := heap.NewMin()
	//for i := 0; i < len(nums); i++ {
	//	sortHeap.Insert(heap.Int(nums[i]))
	//}
	//for i := 0; i < len(nums); i++ {
	//	nums[i] = int(sortHeap.Extract().(heap.Int))
	//}

	for i := (len(nums)-2) / 2; i >= 0; i-- {
		// 建立堆
		filterDown(i, len(nums)-1, nums)
	}
	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		// 不断调整为最大堆
		filterDown(0, i-1, nums)
	}
}

// 堆的建立和调整
func filterDown(current, last int, nums []int)  {
	// child为current的子女位置
	child := 2 * current + 1
	// 暂存子树根结点
	temp  := nums[current]
	// 判断是否到最后结尾
	for child <= last {
		if child < last && nums[child] < nums[child+1] {
			// 让child指向两个孩子中的最大的
			child ++
		}
		// temp的关键码大则不作调整
		if temp >= nums[child] {
			break
		} else {
			// 否则孩子中大者上移
			nums[current] = nums[child]
			current = child
			child = 2 * child + 1
		}
	}
	nums[current] = temp
}