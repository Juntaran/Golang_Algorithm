/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/13 13:59
  */

package Queue

import (
	"sync"
	"fmt"
	"errors"
)

type queueNode struct {				// 队列节点结构
	queueEle 	interface{}			// 元素值
	next 		*queueNode			// next指针
}

type Queue struct {					// 队列结构
	current		*queueNode			// 队列元素
	last 		*queueNode			// 尾指针
	length		int					// 队列长度
	lock 		*sync.Mutex			// 锁
}

// 创建一个队列，返回值是一个Queue
func New() *Queue {
	newQueue := new(Queue)
	newQueue.length = 0
	newQueue.lock = new(sync.Mutex)
	return newQueue
}

// 元素入队
func (queue *Queue) Enqueue(element interface{})  {
	// 先加锁
	queue.lock.Lock()
	defer queue.lock.Unlock()
	{
		// 锁内操作
		if queue.length == 0 {
			queue.current = &queueNode{
				queueEle: 	element,
				next:		nil,
			}
			queue.last = queue.current
		} else {
			newEle := &queueNode{
				queueEle: 	element,
				next:		nil,
			}
			queue.last.next = newEle
			queue.last = newEle
		}
		queue.length ++		// 无论是否原队列为空都长度+1
	}
}

// 元素出队，返回的是队首元素，通过指针改变原队列
func (queue *Queue) Dequeue() (interface{}, error)  {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	{
		if queue.length > 0 {
			element := queue.current.queueEle
			queue.current = queue.current.next
			queue.length --
			return element, nil
		}
		return nil, errors.New("Error: Queue is Empty")
	}
}

// 获取队列长度
func (queue *Queue) LenQueue() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	return queue.length
}

// 判断队列是否为空
func (queue *Queue) IsEmpty() bool {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	return queue.length == 0
}

// 遍历输出队列
func (queue *Queue) PrintQueue() {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	{
		if queue.length == 0 {
			fmt.Println("Queue is Empty")
			return
		}
		var temp = queue.current
		for i := 0; i < queue.length; i++ {
			fmt.Printf("%v ", temp.queueEle)
			temp = temp.next
		}
		fmt.Println()
	}
}

// 返回队列指定位置元素
func (queue *Queue) Pos(index int) (interface{}, error) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	{
		if index >= queue.length || index < 0 {
			return nil, errors.New("Error: Pos Index Out of Range")
		}
		var temp = queue.current
		for i := 0; i < index; i++ {
			temp = temp.next
		}
		return temp.queueEle, nil
	}
}