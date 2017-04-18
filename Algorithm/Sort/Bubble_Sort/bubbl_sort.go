/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 15:28
  */

package Bubble_Sort

func Bubble_Sort(nums []int)  {
	for i := len(nums); ; i-- {
		swap := false
		for j := 1; j <= i; j++ {
			if nums[j-1] > nums[j] {
				swap = true
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
		if swap == false {
			break
		}
	}
}
