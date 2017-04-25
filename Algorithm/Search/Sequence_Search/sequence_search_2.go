/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/20 14:24
  */

package Sequence_Search

import "fmt"

// 哨兵顺序查找
func Sequence_Search_2(nums []int, key int) int {
	if nums[0] == key {
		fmt.Println("Find", key, "in 0")
		return 0
	}
	i := len(nums)-1
	for nums[i] != key && i != 0 {
		i --
	}
	if i == 0 {
		fmt.Println("Not Find")
		return -1
	} else {
		fmt.Println("Find", key, "in", i)
		return i
	}
}
