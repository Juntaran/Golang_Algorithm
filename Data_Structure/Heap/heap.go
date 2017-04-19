package heap

import (
	"sync"
)

type Int int

func (x Int) Less(than Item) bool {
	return x < than.(Int)
}

type Item interface {
	Less(than Item) bool
}

type Heap struct {
	data 	[]Item
	min 	bool
	mutex 	sync.Mutex
}

// 新建一个堆
func New() *Heap {
	newHeap := &Heap{
		data: 	make([]Item, 0),
	}
	return newHeap
}

// 新建一个小顶堆
func NewMin() *Heap {
	newHeap := &Heap{
		data: 	make([]Item, 0),
		min:	true,
	}
	return newHeap
}

// 新建一个大顶堆
func NewMax() *Heap {
	newHeap := &Heap{
		data: 	make([]Item, 0),
		min:	false,
	}
	return newHeap
}

// 堆是否为空
func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}

// 堆大小
func (h *Heap) Size() int {
	return len(h.data)
}

// 获取元素
func (h *Heap) Get(index int) Item {
	return h.data[index]
}

// 插入元素
func (h *Heap) Insert(n Item) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	{
		h.data = append(h.data, n)
		h.siftUp()
		return
	}
}

// 取出元素
func (h *Heap) Extract() (element Item) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	{
		if h.Size() == 0 {
			return
		}
		element = h.data[0]
		last := h.data[h.Size()-1]
		if h.Size() == 1 {
			h.data = nil
			return
		}
		h.data = append([]Item{last}, h.data[1:h.Size()-1]...)
		h.siftDown()
		return
	}
}

// 上移操作
func (h *Heap) siftUp() {
	for i, parent := h.Size()-1, h.Size()-1; i > 0; i = parent {
		parent = i >> 1
		// 如果i位置小于 其父节点，交换
		if h.Less(h.Get(i), h.Get(parent)) {
			h.data[parent], h.data[i] = h.data[i], h.data[parent]
		} else {
			break
		}
	}
}

// 下移操作
func (h *Heap) siftDown() {
	for i, child := 0, 1; i < h.Size() && i<<1+1 < h.Size(); i = child {
		child = i << 1 + 1
		// 两个孩子节点之间交换
		if child+1 <= h.Size()-1 && h.Less(h.Get(child+1), h.Get(child)) {
			child ++
		}
		// 父节点如果小于子节点，交换
		if h.Less(h.Get(i), h.Get(child)) {
			break
		}
		h.data[i], h.data[child] = h.data[child], h.data[i]
	}
}


func (h *Heap) Less(a, b Item) bool {
	if h.min {
		return a.Less(b)
	} else {
		return b.Less(a)
	}
}