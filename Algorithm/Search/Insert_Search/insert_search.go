/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/20 14:29
  */

package Insert_Search

import "fmt"

// 插值查找
func Insert_Search(nums []int, key int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// 差值查找只在middle的取值和二分不同
		middle := left + (right-left) * (key-nums[left]) / (nums[right] - nums[left])
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
