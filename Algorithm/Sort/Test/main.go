/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/19 12:08
  */

package main

import (
	"time"
	"math/rand"
	"fmt"
	"Golang_Algorithm/Algorithm/Sort/Insert_Sort"
	"Golang_Algorithm/Algorithm/Sort/Binary_Insert_Sort"
	"Golang_Algorithm/Algorithm/Sort/Shell_Sort"
	"Golang_Algorithm/Algorithm/Sort/Select_Sort"
	"Golang_Algorithm/Algorithm/Sort/Heap_Sort"
	"Golang_Algorithm/Algorithm/Sort/Bubble_Sort"
	"Golang_Algorithm/Algorithm/Sort/Comb_Sort"
	"Golang_Algorithm/Algorithm/Sort/Shaker_Sort"
	"Golang_Algorithm/Algorithm/Sort/Quick_Sort"
	"Golang_Algorithm/Algorithm/Sort/Merge_Sort"
	"Golang_Algorithm/Algorithm/Sort/Count_Sort"
	"Golang_Algorithm/Algorithm/Sort/Radix_Sort"
)

const Max = 50000		// 随机数生成范围
const Num = 50000		// 随机数生成个数

var start_time  time.Time
var during_time time.Duration

func resetNums(initial, nums []int) []int {
	for i := 0; i < Num; i++ {
		nums[i] = initial[i]
	}
	return nums
}

func main() {
	var initial = make([]int, Num)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Num; i++ {
		initial[i] = rand.Intn(Max)
	}
	fmt.Printf("Initial:\n%v\n", initial)

	var nums = make([]int, Num)
	nums = resetNums(initial, nums)
	//fmt.Printf("%v\n", nums)

	// 直接插入排序
	start_time = time.Now()
	Insert_Sort.InsertSort(nums)
	during_time = time.Since(start_time)
	fmt.Println("Insert Sort: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 二分插入排序
	start_time = time.Now()
	Binary_Insert_Sort.Binary_Insert_Sort(nums)
	during_time = time.Since(start_time)
	fmt.Println("Binary Insert Sort: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 希尔排序
	start_time = time.Now()
	Shell_Sort.Shell_Sort(nums)
	during_time = time.Since(start_time)
	fmt.Println("Shell Sort: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 梳排序
	start_time = time.Now()
	Comb_Sort.Comb_Sort(nums)
	during_time = time.Since(start_time)
	fmt.Println("Comb Sort: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 直接选择排序
	start_time = time.Now()
	Select_Sort.Select_Sort(nums)
	during_time = time.Since(start_time)
	fmt.Println("Select Sort: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 堆排序
	start_time = time.Now()
	Heap_Sort.Heap_Sort(nums)
	during_time = time.Since(start_time)
	fmt.Println("Heap Sort: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 冒泡排序
	start_time = time.Now()
	Bubble_Sort.Bubble_Sort(nums)
	during_time = time.Since(start_time)
	fmt.Println("Bubble Sort: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 鸡尾酒排序
	start_time = time.Now()
	Shaker_Sort.Shaker_Sort(nums)
	during_time = time.Since(start_time)
	fmt.Println("Shaker Sort: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 快速排序递归
	start_time = time.Now()
	Quick_Sort.Quick_Sort_Recurse(nums)
	during_time = time.Since(start_time)
	fmt.Println("Quick Sort Recurse: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 快速排序非递归
	start_time = time.Now()
	Quick_Sort.Quick_Sort_Unrecurse(nums)
	during_time = time.Since(start_time)
	fmt.Println("Quick Sort Unrecurse: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 归并排序
	start_time = time.Now()
	Merge_Sort.Merge_Sort(nums)
	during_time = time.Since(start_time)
	fmt.Println("Merge Sort: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 计数排序
	start_time = time.Now()
	Count_Sort.Count_Sort(nums, Max)
	during_time = time.Since(start_time)
	fmt.Println("Count Sort: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)

	// 基数排序
	start_time = time.Now()
	Radix_Sort.Radix_Sort(nums, Max)
	during_time = time.Since(start_time)
	fmt.Println("Radix Sort: ", during_time)
	fmt.Printf("%v\n", nums)
	nums = resetNums(initial, nums)
}
