/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 15:09
  */

package Shell_Sort

func Shell_Sort(nums []int)  {
	var insertNum int
	increment := (len(nums) + 1) / 2
	for increment > 0 {
		for i := increment; i < len(nums); i++ {
			insertNum = nums[i]
			j := i
			for j >= increment && insertNum < nums[j-increment] {
				nums[j] = nums[j-increment]
				j -= increment
			}
			// 插入
			nums[j] = insertNum
		}
		// 缩小增量
		increment /= 2
	}
}