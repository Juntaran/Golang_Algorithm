/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/20 14:22
  */

package Sequence_Search

func Sequence_Search(nums []int, key int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			return i
		}
	}
	return -1
}