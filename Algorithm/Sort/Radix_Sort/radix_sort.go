/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/19 11:13
  */

package Radix_Sort

// 取一个数字从低到高第pos位的位数
func getNumInPos(num, pos int) int {
	temp := 1
	for i := 0; i < pos - 1; i++ {
		temp *= 10
	}
	return (num / temp) % 10
}

// 取一个数字的位数
func getDigt(num int) int {
	n := 0
	for num > 0 {
		n ++
		num /= 10
	}
	return n
}

// 基数排序
func Radix_Sort(nums []int, max int)  {
	// 获取待排序数的最大数位数
	digitRandom := getDigt(max)
	// 数组对应0-9序列
	//length := len(nums) + 1
	//var radixArray [10][len(nums) + 1]int
	radixArray := [10]([]int){}
	// index为0处记录这组数据的个数
	for i := 0; i < 10; i++ {
		radixArray[i] = make([]int, len(nums)+1)
		radixArray[i][0] = 0
	}
	for pos := 1; pos <= digitRandom; pos++ {
		// 分配过程
		for i := 0; i < len(nums); i++ {
			num := getNumInPos(nums[i], pos)
			radixArray[num][0] ++
			index := radixArray[num][0]
			radixArray[num][index] = nums[i]
		}
		// 收集
		j := 0
		for i := 0; i < 10; i++ {
			for k := 1; k <= radixArray[i][0]; k++ {
				nums[j] = radixArray[i][k]
				j ++
			}
			// 复位
			radixArray[i][0] = 0
		}
	}
}