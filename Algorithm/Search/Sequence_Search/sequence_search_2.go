/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/20 14:24
  */

package Sequence_Search

// 哨兵顺序查找
func Sequence_Search_2(nums []int, key int) int {
	if nums[0] == key {
		return 0
	}
	i := len(nums)
	for nums[i] != key && i != 0 {
		i --
	}
	if i == 0 {
		return -1
	} else {
		return i
	}
}
