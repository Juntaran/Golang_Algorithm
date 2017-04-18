/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/17 15:33
  */

package Binary_Tree

import (
	"errors"
	"fmt"
)

// 树节点结构
type TreeNode struct {
	Value 		int
	parent 		*TreeNode
	lchild 		*TreeNode
	rchild		*TreeNode
}

// 树结构
type Tree struct {
	Head 		*TreeNode
	Size 		int
}


// 创建一个节点
func NewNode(value int) *TreeNode {
	newNode := &TreeNode{
		Value:  	value,
	}
	return newNode
}

// 对比节点值
func (node *TreeNode) Compare(m *TreeNode) int {
	if node.Value < m.Value {
		return -1
	} else if node.Value > m.Value {
		return 1
	} else {
		return 0
	}
}

// 创建一棵树
func NewTree(node *TreeNode) *Tree {
	if node == nil {
		return &Tree{}
	}
	return &Tree{
		Head: 	node,
		Size: 	1,
	}
}

// 插入元素
func (tree *Tree)Insert(value int) {
	node := NewNode(value)				// node是待插入节点
	if tree.Size == 0 {
		tree.Head = node
		tree.Size ++
		return
	}
	head := tree.Head
	for {
		if node.Compare(head) == -1 {
			if head.lchild == nil {
				head.lchild = node
				node.parent = head
				break
			} else {
				head = head.lchild
			}
		} else {
			if head.rchild == nil {
				head.rchild = node
				node.parent = head
				break
			} else {
				head = head.rchild
			}
		}
	}
	tree.Size ++
}

// 搜索元素
func (tree *Tree) Search(value int) (*TreeNode, error) {
	head := tree.Head			// 待搜索的树
	node := NewNode(value)		// 待搜索元素

	for head != nil {
		switch head.Compare(node) {
		case -1:
			head = head.rchild
		case 1:
			head = head.lchild
		case 0:
			return head, nil
		default:
			err := errors.New("Node not Found")
			return nil, err
		}
	}
	err := errors.New("Error: Search Node not Found")
	return nil, err
}

// 删除元素
func (tree *Tree) Delete(value int) error {
	count := 0
Search:
	_, err := tree.Search(value)
	//fmt.Println(ret.Value)
	// 该元素不存在
	if err != nil {		// err != nil 表示没有搜索到这个元素
		if count == 0 {
			return errors.New("Error: Delete not Found")
		}
		return nil
	}

	var parent *TreeNode
	head := tree.Head
	node := NewNode(value)
	for node != nil {
		switch node.Compare(head) {
		case -1:
			parent = head
			head = head.lchild
		case 1:
			parent = head
			head = head.rchild
		case 0:
			// 找到的节点是叶子节点
			if head.lchild == nil && head.rchild == nil {
				if parent.lchild == head {
					parent.lchild = nil
				} else {
					parent.rchild = nil
				}
				count ++
				tree.Size --
				goto Search
			}
			// 搜索到该节点，且该节点左孩子不为空
			if head.lchild != nil {
				right := head.rchild
				temp := head.lchild.rchild
				head.Value = head.lchild.Value
				head.lchild = head.lchild.lchild
				if head.lchild != nil {
					head.lchild.parent = head
				}
				head.rchild = temp
				if head.rchild != nil {
					head.rchild.parent = head
				}
				// 如果搜索到的节点有右孩子
				if right != nil {
					// 循环，一直找左子树的最右节点
					tempHead := head
					for tempHead.rchild != nil {
						tempHead = tempHead.rchild
					}
					tempHead.rchild = right
					right.parent = tempHead
				}
				tree.Size --
				count ++
				goto Search
			}
			// 搜索到该节点，且该节点没有左孩子，只有右孩子且不为空
			if head.rchild != nil {
				head.Value = head.rchild.Value
				head.lchild = head.rchild.lchild
				if head.lchild != nil {
					head.lchild.parent = head
				}
				head.rchild = head.rchild.rchild
				if head.rchild != nil {
					head.rchild.parent = head
				}
				tree.Size --
				count ++
				goto Search
			}
			// 如果这棵树只有一个节点，就是待删除节点
			if parent == nil {
				tree.Head = nil
				tree.Size --
				count ++
				return nil
			}

		}
	}
	return errors.New("Error: Delete not Found")
}

// 树的遍历
// 前序遍历
func (tree *Tree)PreOrderTraverse()  {
	fmt.Println("Start PrenOrderTraverse")
	preOrder(tree.Head)
	fmt.Println()
}
func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%4d ", root.Value)
	preOrder(root.lchild)
	preOrder(root.rchild)
	return
}

// 中序遍历
func (tree *Tree)InOrderTraverse()  {
	fmt.Println("Start InOrderTraverse")
	inOrder(tree.Head)
	fmt.Println()
}
func inOrder(root *TreeNode)  {
	if root == nil {
		return
	}
	inOrder(root.lchild)
	fmt.Printf("%4d ", root.Value)
	inOrder(root.rchild)
	return
}

// 后序遍历
func (tree *Tree)PostOrderTraverse()  {
	fmt.Println("Start PostOrderTraverse")
	postOrder(tree.Head)
	fmt.Println()
}
func postOrder(root *TreeNode)  {
	if root == nil {
		return
	}
	postOrder(root.lchild)
	postOrder(root.rchild)
	fmt.Printf("%4d ", root.Value)
	return
}

// 层序遍历
func (tree *Tree)LevelTraverse()  {
	fmt.Println("Start LevelTraverse")
	if tree.Head == nil {
		return
	}
	root := tree.Head
	var treeQueue []*TreeNode
	treeQueue = append(treeQueue, root)
	var current int = 0
	var last int = 1
	for current < len(treeQueue) {
		last = len(treeQueue)
		for current < last {
			fmt.Printf("%4d ", treeQueue[current].Value)
			if treeQueue[current].lchild != nil {
				treeQueue = append(treeQueue, treeQueue[current].lchild)
			}
			if treeQueue[current].rchild != nil {
				treeQueue = append(treeQueue, treeQueue[current].rchild)
			}
			current ++
		}
		fmt.Println()
	}
	fmt.Println()
}
