/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 14:40
  */

package Insert_Sort

func InsertSort(nums []int)  {
	var insertNum int
	for i := 1; i < len(nums); i++ {
		insertNum = nums[i]		// 待插入的元素
		j := i - 1
		for j >= 0 && insertNum < nums[j] {
			nums[j+1] = nums[j]
			j --
		}
		nums[j+1] = insertNum
	}
}