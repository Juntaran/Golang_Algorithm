/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 17:15
  */

package Quick_Sort

import "fmt"

func Quick_Sort_Recurse_Single(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	flag := nums[0]
	left, right := 0, len(nums) - 1

	fmt.Println(nums, "\tleft:", left, "right:", right, "flag:", flag)

	for i := 1; i <= right; {
		if nums[i] > flag {
			nums[i], nums[right] = nums[right], nums[i]
			right --
			fmt.Println(nums, "\tleft:", left, "right:", right, "flag:", flag)
		} else {
			nums[i], nums[left] = nums[left], nums[i]
			left ++
			i ++
			fmt.Println(nums, "\tleft:", left, "right:", right, "flag:", flag)
		}
	}

	Quick_Sort_Recurse_Single(nums[:left])
	Quick_Sort_Recurse_Single(nums[left+1:])
}
