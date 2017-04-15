/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/15 09:06
  */

package Hash_Table

import (
	"Golang_Algorithm/Data_Structure/Link_List"
	"fmt"
	"math"
	"errors"
)

// 设置哈希表桶个数，如果达到最大会resize，翻倍桶数量
const fillRate int = 10

// 哈希表结构
type HashTable struct {
	table 		[]*Link_List.List		// 底层结构
	size		int						// 当前哈希表中元素数
	capacity 	int						// 最大容量，超过会resize
}

// 子结构把kv加入哈希表，这是存储在每个bucket的链表的值
type item struct {
	key 		string
	value 		interface{}
}

// 使用的哈希算法
func hashCode(s string) int {
	hash := int32(0)
	for i := 0; i < len(s); i++ {
		hash = int32(hash << 5-hash) + int32(s[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}

// 利用string类型的key来寻找bucket，也就是第几个链表
func (hashTable *HashTable)getIndex(key string, length int) int {
	return hashCode(key) % length
}

// 在指定的bucket内寻找指定的key，返回*tableNode
func (hashTable *HashTable) find(index int, key string) (*item, error) {
	list := hashTable.table[index]
	if list == nil {
		return nil, errors.New("Error: Not Found")
	}
	var val *item
	list.DoEach(func(node Link_List.ListNode) {
		// 类型断言 如果传入的值为*hashNode类型，则获取其值，否则panic
		if node.Value.(*item).key == key {
			val = node.Value.(*item)
		}
	})
	if val == nil {
		return nil, errors.New("Error: Not Found")
	}
	return val, nil
}

// 创建一个哈希表
func New(size int) *HashTable {
	newHashTable := &HashTable {
		table: 		make([]*Link_List.List, size),
		size: 		0,
		capacity: 	fillRate * size,
	}
	return newHashTable
}

// 返回一个默认哈希表，默认大小为128
func NewDefault() *HashTable {
	return New(128)
}

// 让哈希表变大，每个桶里的元素就变少了，用更多内存降低碰撞
func (hashTable *HashTable) resizeTable() {
	newHashTable := New(2 * len(hashTable.table))
	for _, list := range hashTable.table {
		if list != nil {
			for _, itemi := range list.Items() {
				// ok的作用是如果类型断言失败，那么parsed仍然存在，不过为零值
				if parsed, ok := (*itemi).(item); ok {
					newHashTable.Insert(parsed.key, parsed.value)
				} else {
					fmt.Println("Failed to Parse item in resize")
				}
			}
		}
	}
	hashTable = newHashTable
}

// 插入一个键值对
func (hashTable *HashTable) Insert(key string, value interface{}) {
	index := hashTable.getIndex(key, len(hashTable.table))
	// 如果当前桶为空
	if hashTable.table[index] == nil {
		hashTable.table[index] = Link_List.New()
	}
	item := &item{
		key: 	key,
		value: 	value,
	}
	ret, err := hashTable.find(index, key)
	if err != nil {
		// 这个bucket内不存在这个key
		hashTable.table[index].Append(item)
		hashTable.size ++
	} else {
		// key存在，覆盖
		ret.value = value
	}
}

// 删除一个键值对
func (hashTable *HashTable) Delete(key string) error {
	index := hashTable.getIndex(key, len(hashTable.table))
	list := hashTable.table[index]
	var val *item

	list.DoEach(func(node Link_List.ListNode) {
		if node.Value.(*item).key == key {
			val = node.Value.(*item)
		}
	})

	if val == nil {
		return errors.New("Error: Delete Not Found")
	}
	hashTable.size --
	return list.Remove(val)
}

// 判断一个key是否存在
func (hashTable *HashTable) ContainsKey(key string) bool {
	index := hashTable.getIndex(key, len(hashTable.table))
	item, _ := hashTable.find(index, key)
	if item == nil {
		return false
	}
	return true
}


// 根据key取value
func (hashTable *HashTable) GetValue(key string) (interface{}, error) {
	index := hashTable.getIndex(key, len(hashTable.table))
	item, err := hashTable.find(index, key)
	if item == nil {
		return "", errors.New("Error: GetValue Not Found")
	}
	return item.value, err
}

// 返回哈希表大小
func (hashTable *HashTable) ReturnSize() int {
	return hashTable.size
}