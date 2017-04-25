/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/25 15:19
  */

package main

import (
	"time"
	"math/rand"
	"fmt"
	"Golang_Algorithm/Algorithm/Sort/Quick_Sort"
	"Golang_Algorithm/Algorithm/Search/Sequence_Search"
	"Golang_Algorithm/Algorithm/Search/Binary_Search"
	"Golang_Algorithm/Algorithm/Search/Insert_Search"
	"Golang_Algorithm/Algorithm/Search/Fibonacci_Search"
	"Golang_Algorithm/Algorithm/Search/Binary_Tree_Search"
	"Golang_Algorithm/Algorithm/Search/Hash_Search"
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
	// initial是无序数组
	var initial = make([]int, Num)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Num; i++ {
		initial[i] = rand.Intn(Max)
	}
	fmt.Printf("Initial:\n%v\n", initial)
	// nums是有序数组
	var nums = make([]int, Num)
	nums = resetNums(initial, nums)
	Quick_Sort.Quick_Sort_Recurse(nums)
	fmt.Printf("Nums:\n%v\n", nums)

	var key = initial[Num / 2]

	// 顺序查找
	fmt.Println("Sequence Search:")
	start_time = time.Now()
	Sequence_Search.Sequence_Search(initial, key)
	during_time = time.Since(start_time)
	fmt.Println("Sequence Search Time:", during_time)

	// 哨兵顺序查找
	fmt.Println("Sequence Search2:")
	start_time = time.Now()
	Sequence_Search.Sequence_Search_2(initial, key)
	during_time = time.Since(start_time)
	fmt.Println("Sequence Search2 Time:", during_time)

	// 二分查找
	fmt.Println("Binary Search:")
	start_time = time.Now()
	Binary_Search.Binary_Search(nums, key)
	during_time = time.Since(start_time)
	fmt.Println("Binary Search Time:", during_time)

	// 插值查找
	fmt.Println("Insert Search:")
	start_time = time.Now()
	Insert_Search.Insert_Search(nums, key)
	during_time = time.Since(start_time)
	fmt.Println("Insert Search Time:", during_time)

	// 斐波那契查找
	fmt.Println("Fibonacci Search:")
	start_time = time.Now()
	Fibonacci_Search.Fibonacci_Search(nums, key)
	during_time = time.Since(start_time)
	fmt.Println("Fibonacci Search Time:", during_time)

	// 二叉搜索树
	fmt.Println("Binary Tree Search")
	start_time = time.Now()
	Binary_Tree_Search.Binary_Tree_Search(nums, key)
	during_time = time.Since(start_time)
	fmt.Println("Binary Tree Search Time:", during_time)

	// 哈希表查找
	fmt.Println("Hash Table Search")
	start_time = time.Now()
	Hash_Search.Hash_Search(nums, key)
	during_time = time.Since(start_time)
	fmt.Println("Hash Table Search Time:", during_time)
}
