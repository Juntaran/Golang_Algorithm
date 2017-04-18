/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 17:07
  */

package Quick_Sort

import "Golang_Algorithm/Data_Structure/Stack"

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
	sort_stack := Stack.New()
	left  := 0
	right := len(nums) - 1
	if left < right {
		sort_stack.Push(left)
		sort_stack.Push(right)
		for sort_stack.IsEmpty() == false {
			j := sort_stack.Top().(int)
			sort_stack.Pop()
			i := sort_stack.Top().(int)
			sort_stack.Pop()
			k := partition(nums, i, j)
			if i < k - 1 {
				sort_stack.Push(i)
				sort_stack.Push(k-1)
			}
			if k + 1 < j {
				sort_stack.Push(k+1)
				sort_stack.Push(j)
			}
		}
	}
}