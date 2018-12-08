/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 17:15
  */

package Quick_Sort

func Quick_Sort_Recurse(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left < right {
		i, j := left, right
		flag := nums[left]
		// 从右向左找第一个比 flag 小的数
		for i < j {
			if nums[j] >= flag {
				j --
			} else {
				break
			}
		}
		nums[i] = nums[j]
		// 从左向右找第一个比 flag 大的数
		for i < j {
			if nums[i] < flag {
				i ++
			} else {
				break
			}
		}
		nums[j] = nums[i]
		nums[i] = flag
		quickSort(nums, left, i-1)
		quickSort(nums, i+1, right)
	}
}