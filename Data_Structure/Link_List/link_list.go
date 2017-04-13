/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/13 16:36
  */

package Link_List

import (
	"sync"
	"errors"
	"fmt"
)

type ListNode struct {			// 链表元素结构
	value	interface{}
	next 	*ListNode
	prev 	*ListNode
}

type List struct {				// 链表结构
	head 	*ListNode
	tail 	*ListNode
	length 	int
	lock 	*sync.Mutex
}

// 创建链表
func New() *List {
	newList := new(List)
	newList.length = 0
	newList.lock = new(sync.Mutex)
	return newList
}

// 在链表头插入
func (list *List) Prepend(value interface{}) {
	list.lock.Lock()
	defer list.lock.Unlock()
	{
		node := new(ListNode)
		node.value = value
		if list.length == 0 {
			list.head = node
			list.tail = node
		} else {
			list.head.prev = node
			node.next = list.head
			list.head = node
		}
		list.length ++
	}
}

// 在链表尾插入
func (list *List) Append(value interface{}) {
	list.lock.Lock()
	defer list.lock.Unlock()
	{
		node := new(ListNode)
		node.value = value

		if list.length == 0 {
			list.head = node
			list.tail = node
		} else {
			list.tail.next = node
			node.prev = list.tail
			node.next = nil
			list.tail = node
		}
		list.length ++
	}
}

// 在index处插入链表
func (list *List) Insert(value interface{}, index int) error {
	if index <= 0 || index > list.length {
		return errors.New("Error: Insert Index Out of Range")
	}
	list.lock.Lock()
	defer list.lock.Unlock()
	{
		node := new(ListNode)
		node.value = value
		if list.length == 0 {
			list.head = node
			list.tail = node
			//node.prev = nil
			//node.next = nil
		} else if index == 1{
			node.next = list.head
			list.head.prev = node
			list.head = node
		} else {
			temp := list.head
			for i := 1; i < index; i++ {
				temp = temp.next
			}
			temp.prev.next = node
			node.prev = temp.prev
			node.next = temp
			temp.prev = node
		}
		list.length ++
	}
	return nil
}

// 删除index位置的元素
func (list *List) DeleteIndex(index int) error {
	list.lock.Lock()
	defer list.lock.Unlock()
	{
		if list.length <= 0 {
			return errors.New("Error: List is Empty")
		}
		if index <= 0 || index > list.length {
			return errors.New("Error: Delete Index Out of Range")
		}
		temp := list.head
		for i := 1; i < index; i++ {
			temp = temp.next
		}
		temp.prev.next = temp.next
		temp.next.prev = temp.prev
		list.length --
		return nil
	}
}

// 删除所有值为value的元素
func (list *List) DeleteValue(value interface{}) (int, error) {
	list.lock.Lock()
	defer list.lock.Unlock()
	{
		count := 0
Here:
		if count == 0 && list.length <= 0 {
			return 0, errors.New("Error: Empty List")
		} else if list.length <= 0 {
			// 第一个元素为value，且只有一个元素
			return count, nil
		}
		// 链表头的值恰好为value
		if list.head.value == value {
			list.head = list.head.next
			list.length --
			count ++
			goto Here
		}
		for node := list.head; node.next != nil; node = node.next {
			if node.value == value {
				node.prev.next = node.next
				node.next.prev = node.prev
				list.length --
				count ++
			}
		}
		if list.tail.value == value {
			count ++
			list.tail.prev.next = nil
			list.length --
		}
		return count, nil
	}
}

// 获取链表长度
func (list *List) LenList() int {
	list.lock.Lock()
	defer list.lock.Unlock()
	return list.length
}

// 判断链表是否为空
func (list *List) IsEmpty() bool {
	list.lock.Lock()
	defer list.lock.Unlock()
	return list.length == 0
}

// 清空链表
func (list *List) Clear(){
	list.lock.Lock()
	defer list.lock.Unlock()
	{
		list.length = 0
		list.head = nil
		list.tail = nil
	}
}

// 附加链表k
func (list *List) Connect(k *List) {
	list.lock.Lock()
	defer list.lock.Unlock()
	{
		if list.length <= 0 {
			list = k
			list.length = k.length
			return
		}
		if k.length <= 0 {
			list = list
			return
		}
		list.tail.next, k.head.prev = k.head, list.tail
		list.tail = k.tail
		list.length += k.length
	}
}

// 获取index位置的元素值
func (list *List) GetIndex(index int) (interface{}, error) {
	list.lock.Lock()
	defer list.lock.Unlock()
	{
		if index <= 0 || index > list.length {
			return nil, errors.New("Error: GetIndex Index Out of Range")
		}
		node := list.head
		for i := 1; i < index; i++ {
			node = node.next
		}
		return node.value, nil
	}
}

// 寻找链表中值为value元素的位置，返回一个切片和总数
func (list *List) FindIndex(value interface{}) ([]int, int) {
	list.lock.Lock()
	defer list.lock.Unlock()
	{
		if list.length <= 0 {
			return nil, 0
		}
		var result []int
		count := 0
		index := 1
		for node := list.head; node != nil; node = node.next {
			if node.value == value {
				count ++
				result = append(result, index)
			}
			index ++
		}
		return result, count
	}
}

// 替换index位置的元素
func (list *List) ChangeList(index int, value interface{}) error {
	if index <= 0 || index > list.length {
		return errors.New("Error: ChangeList Index Out of Range")
	}
	list.DeleteIndex(index)
	list.Insert(value, index)
	return nil
}

// 遍历输出切片
func (list *List) PrintList() {
	list.lock.Lock()
	defer list.lock.Unlock()
	{
		if list.length == 0 {
			fmt.Println("List is Empty")
			return
		}
		for node := list.head; node != nil; node = node.next {
			fmt.Printf("%v ", node.value)
		}
		fmt.Println()
	}
}