/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2019/5/14 02:37
  */

package Quick_Sort

/*
	双路快排
	从左向右找到大于 flag 的元素
	从右向左找到小于 flag 的元素
	判断是否越界，交换元素，left++, right--
*/

// TODO：双路快排存在一些问题	https://blog.csdn.net/u012279631/article/details/80717445
func Quick_Sort_Recurse_Double(nums []int)  {
	length := len(nums)
	if length <= 1 {
		return
	}

	flag := nums[0]
	left, right := 0, length - 1

	for {
		for left <= right && nums[left] < flag {
			left ++
		}
		for right > left && nums[right] > flag {
			right --
		}
		if left > right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left ++
		right --
	}

	Quick_Sort_Recurse_Double(nums[:left])
	Quick_Sort_Recurse_Double(nums[left+1:])
}
