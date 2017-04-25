/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/25 14:32
  */

package Hash_Search

import (
	"fmt"
)

const numkey int = -32768

type hashTable struct {
	element 	[]int		// 数据元素存储基址，动态分配切片
	count 		int			// 当前数据元素个数
}

// 创建一个哈希表
func newTable() *hashTable {
	newHashTable := &hashTable {
		element: 	make([]int, 5),
		count: 		0,
	}
	return newHashTable
}

// 初始化哈希表
func initHashTable(table *hashTable, size int)  {
	table.count = size
	table.element = make([]int, size)
	for i := 0; i < size; i++ {
		table.element[i] = numkey
	}
}

// 哈希函数
func hash(key int, size int) int {
	// 除留余数法
	return key % size
}

// 插入数据
func insertHash(table *hashTable, key int)  {
	// 获取哈希地址
	addr := hash(key, table.count)
	// 如果不为空，则冲突
	for table.element[addr] != numkey {
		// 开放定址法的线性探测
		addr = (addr + 1) % table.count
	}
	// 找到空位后插入
	table.element[addr] = key
}

// 在哈希表中查找
func search(table *hashTable, key int) (bool, int) {
	// 求哈希地址
	addr := hash(key, table.count)
	// 冲突
	for table.element[addr] != key {
		// 线性探测
		addr = (addr + 1) % table.count
		if table.element[addr] == numkey || addr == hash(key, table.count) {
			return false, -1
		}
	}
	return true, addr
}

func Hash_Search(nums []int, key int)  {
	table := newTable()
	initHashTable(table, len(nums))
	for i := 0; i < len(nums); i++ {
		insertHash(table, nums[i])
	}
	ret, addr := search(table, key)
	if ret == false {
		fmt.Println("Not Find")
	} else {
		fmt.Println("Find", key, "in", addr)
	}
}