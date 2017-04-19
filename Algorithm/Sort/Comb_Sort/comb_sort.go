/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/19 12:05
  */

package Comb_Sort

func Comb_Sort(nums []int)  {
	gap := len(nums)
	for {
		if gap > 1 {
			gap = gap * 100 / 124
		}
		for i := 0; (i+gap) < len(nums); i++ {
			if nums[i] > nums[i+gap] {
				nums[i], nums[i+gap] = nums[i+gap], nums[i]
			}
		}
		if gap == 1 {
			break
		}
	}
}
