/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 17:07
  */

package Quick_Sort

//import "Golang_Algorithm/Data_Structure/Stack"

func partition(nums []int, left int, right int) int {
	// 选第一个元素作为枢纽元
	flag := nums[left]
	for left < right {
		// 从右向左找到第一个小于flag的元素，放到left位置
		for left < right && nums[right] >= flag {
			right --
		}
		nums[left] = nums[right]

		// 从左向右找到第一个大于flag的元素，放到right位置
		for left < right && nums[left] <= flag {
			left ++
		}
		nums[right] = nums[left]
	}
	nums[left] = flag
	return left
}

func Quick_Sort_Unrecurse(nums []int)  {

/***********************方法1***************************/
	//sort_stack := Stack.New()
	//left  := 0
	//right := len(nums) - 1
	//if left < right {
	//	sort_stack.Push(left)
	//	sort_stack.Push(right)
	//	for sort_stack.IsEmpty() == false {
	//		j := sort_stack.Top().(int)
	//		sort_stack.Pop()
	//		i := sort_stack.Top().(int)
	//		sort_stack.Pop()
	//		k := partition(nums, i, j)
	//		if i < k - 1 {
	//			sort_stack.Push(i)
	//			sort_stack.Push(k-1)
	//		}
	//		if k + 1 < j {
	//			sort_stack.Push(k+1)
	//			sort_stack.Push(j)
	//		}
	//	}
	//}

/***********************方法2***************************/
	//// 栈中保存下次需要排序的子数组的开始位置和结束位置
	//sort_stack := Stack.New()
	//sort_stack.Push(0)
	//sort_stack.Push(len(nums) - 1)
	//
	//// 栈非空
	//for sort_stack.IsEmpty() == false {
	//	high := sort_stack.Top().(int)
	//	sort_stack.Pop()
	//	low  := sort_stack.Top().(int)
	//	sort_stack.Pop()
	//	middle := partition(nums, low, high)
	//
	//	// 右边子数组入栈
	//	if (middle + 1) < high {
	//		sort_stack.Push(middle + 1)
	//		sort_stack.Push(high)
	//	}
	//	// 左边数组入栈
	//	if low < (middle - 1) {
	//		sort_stack.Push(low)
	//		sort_stack.Push(middle - 1)
	//	}
	//}

/***********************方法3***************************/
	// 递归层数设置，栈中保存下次需要排序的子数组的开始位置和结束位置
	var sort_stack = make([]int, len(nums) * 2)
	top := -1
	top ++
	sort_stack[top] = 0
	top ++
	sort_stack[top] = len(nums) - 1

	// 栈非空
	for top > 0 {
		high := sort_stack[top]
		top --
		low  := sort_stack[top]
		top --
		middle := partition(nums, low, high)

		// 右边子数组入栈
		if middle + 1 < high {
			top ++
			sort_stack[top] = middle + 1
			top ++
			sort_stack[top] = high
		}
		// 左边子数组入栈
		if middle - 1 > low {
			top ++
			sort_stack[top] = low
			top ++
			sort_stack[top] = middle - 1
		}
	}

}