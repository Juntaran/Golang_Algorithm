/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 14:36
  */

package Binary_Insert_Sort

func Binary_Insert_Sort(nums []int)  {
	var insertNum int
	var middle int
	for i := 0; i < len(nums); i++ {
		insertNum = nums[i]
		left  := 0				// 已排序数组左边界
		right := i - 1			// 已排序数组右边界
		for left <= right {		// 二分法寻找插入位置
			middle = (left + right) / 2
			if insertNum > nums[middle] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
		for j := i; i > left ; j-- {
			nums[j] = nums[j-1]
		}
		nums[left] = insertNum
	}
}