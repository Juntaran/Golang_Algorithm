/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 17:15
  */

package Quick_Sort

func Quick_Sort_Recurse(nums []int)  {
	left  := 0
	right := len(nums) - 1
	sort(nums, left, right)
}

func sort(nums []int, left int, right int)  {
	if left < right {
		i := left
		j := right
		flag := nums[left]

		for i < j {
			// 从右向左找第一个比基数小的元素
			for i < j && nums[j] >= flag {
				j --
			}
			nums[i] = nums[j]

			// 从左向右找第一个比基数大的元素
			for i < j && nums[i] < flag {
				i ++
			}
			nums[j] = nums[i]
		}
		nums[i] = flag
		sort(nums, left, i-1)
		sort(nums, i+1,  right)
	}
}