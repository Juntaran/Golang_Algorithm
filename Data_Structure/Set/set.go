/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/5 00:43
  */

package Set

import (
	"sync"
	"fmt"
	"errors"
)

type Set struct {
	set 	map[interface{}]bool
	lock 	sync.Mutex
}

// 创建一个集合
func New() *Set {
	return &Set{
		set: 	map[interface{}]bool{},
	}
}

// 添加元素
func (set *Set) Add(item interface{})  {
	set.lock.Lock()
	defer set.lock.Unlock()
	{
		set.set[item] = true
	}
}

// 删除元素
func (set *Set) Delete(item interface{}) error {
	if set.IsExist(item) == false {
		return errors.New("Error: Delete Not Find")
	}
	set.lock.Lock()
	defer set.lock.Unlock()
	{
		delete(set.set, item)
		return nil
	}
}

// 判断该元素是否存在
func (set *Set) IsExist(item interface{}) bool {
	set.lock.Lock()
	defer set.lock.Unlock()
	{
		_, ok := set.set[item]
		return ok
	}
}

// 集合大小
func (set *Set) LenSet() int {
	return len(set.listSet())
}

// 判断集合是否为空
func (set *Set) IsEmpty() bool {
	if set.LenSet() == 0 {
		return true
	}
	return false
}

// 输出集合中所有元素
func (set *Set) PrintSet() {
	setList := set.listSet()
	fmt.Println(setList)
}

// 集合元素存入切片
func (set *Set) listSet() []interface{} {
	set.lock.Lock()
	defer set.lock.Unlock()
	{
		list := []interface{}{}
		for item := range set.set {
			list = append(list, item)
		}
		return list
	}
}