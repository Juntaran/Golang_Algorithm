/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 17:26
  */

package Merge_Sort

// 合并数组
func merge(sourceArr []int, tempArr []int, left, middle, right int)  {
	i := left
	j := middle + 1
	k := left

	for i != middle + 1 && j != right + 1 {
		if sourceArr[i] >= sourceArr[j] {
			tempArr[k] = sourceArr[j]
			k ++
			j ++
		} else {
			tempArr[k] = sourceArr[i]
			k ++
			i ++
		}
	}
	for i != middle + 1 {
		tempArr[k] = sourceArr[i]
		k ++
		i ++
	}
	for j != right + 1 {
		tempArr[k] = sourceArr[j]
		k ++
		j ++
	}
	for i = left; i <= right; i++ {
		sourceArr[i] = tempArr[i]
	}
}

// 内部递归
func recurse(sourceArr, tempArr []int, left, right int)  {
	var middle int
	if left < right {
		middle = (left + right) / 2
		recurse(sourceArr, tempArr, left, middle)
		recurse(sourceArr, tempArr, middle+1, right)
		merge(sourceArr, tempArr, left, middle, right)
	}
}

func Merge_Sort(nums []int)  {
	left  := 0
	right := len(nums) - 1
	var tempArr = make([]int, len(nums))
	recurse(nums, tempArr, left, right)
}