/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 15:35
  */

package Shaker_Sort

// 鸡尾酒排序，也就是双向冒泡
func Shaker_Sort(nums []int)  {
	left  := 0
	right := len(nums) - 1
	var flag int			// 左右有序标志位

	for left < right {
		// 从左向右冒泡
		for i := left; i < right; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				flag = i
			}
		}
		right = flag
		// 从右向左冒泡
		for j := right; j > left; j-- {
			if nums[j] < nums[j-1] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
				flag = j
			}
		}
		left = flag
	}
}
