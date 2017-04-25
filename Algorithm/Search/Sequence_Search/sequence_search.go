/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/20 14:22
  */

package Sequence_Search

import "fmt"

func Sequence_Search(nums []int, key int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			fmt.Println("Find", key, "in", i)
			return i
		}
	}
	fmt.Println("Not Find")
	return -1
}