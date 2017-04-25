/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/20 14:33
  */

package Fibonacci_Search

import "fmt"

// 斐波那契查找
func Fibonacci_Search(nums []int, key int) int {
	// 斐波那契数列
	var Fibonacci [1000]int
	Fibonacci[0] = 0
	Fibonacci[1] = 1
	for i := 2; i < 1000; i++ {
		Fibonacci[i] = Fibonacci[i-1] + Fibonacci[i-2]
	}
	left, right := 0, len(nums)-1
	k := 0
	// 计算n位于斐波那契数列的位置
	for len(nums) > Fibonacci[k] - 1 {
		k ++
	}
	// 扩展nums到Fibonacci[k]-1的长度
	var temp = make([]int, Fibonacci[k]-1)
	for i := 0; i < len(nums); i++ {
		temp[i] = nums[i]
	}
	// 填充剩余部分
	for i := len(nums); i < Fibonacci[k]-1; i++ {
		temp[i] = nums[len(nums) - 1]
	}
	for left <= right {
		middle := left + Fibonacci[k-1] - 1
		if key < temp[middle] {
			right = middle - 1
			// 斐波那契数列下标-1
			k = k - 1
		} else if key > temp[middle] {
			left = middle + 1
			// 斐波那契数列下标-2
			k = k - 2
		} else {
			if middle < len(nums) {
				fmt.Println("Find", key, "in", middle)
				return middle
			} else {
				fmt.Println("Find", key, "in", len(nums)-1)
				return len(nums) - 1
			}
		}
	}
	fmt.Println("Not Find")
	return -1
}
