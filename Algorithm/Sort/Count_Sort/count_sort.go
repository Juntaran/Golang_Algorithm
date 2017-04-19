/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/19 11:09
  */

package Count_Sort

func Count_Sort(nums []int, max int)  {
	// max是需要排序的数的数值的最大值
	markArray := make([]int, max)
	for i := 0; i < len(nums); i++ {
		markArray[nums[i]] ++			// 统计每个元素出现的次数
	}
	k := 0
	for i := 0; i < max; i++ {
		if markArray[i] != 0 {
			for j := 0; j < markArray[i]; j++ {
				nums[k] = i
				k ++
			}
		}
	}
}
