/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/20 14:15
  */

package Binary_Search

import "fmt"

// 二分查找首先要数组有序
func Binary_Search(nums []int, key int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (right - left) >> 1 + left
		if key < nums[middle] {
			right = middle - 1
		} else if key > nums[middle] {
			left = middle + 1
		} else {
			fmt.Println("Find", key, "in", middle)
			return middle
		}
	}
	fmt.Println("Not Find")
	return -1
}