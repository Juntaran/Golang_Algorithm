/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/1/30 14:10
  */

package Merkle_Tree

import (
	"hash"
	"errors"
)

// Merkle Tree 节点结构
type MerkleNode struct {
	HashValue	[]byte
	lchild 		*MerkleNode
	rchild		*MerkleNode
}

// Merkle Tree 树结构
type MerkleTree struct {
	Size 		uint64					// 节点数
	Nodes 		[]MerkleNode			// 所有 node
	Level		[][]MerkleNode			// 每层的 node
}

// 创建一个节点
func NewMerkleNode(h hash.Hash, block []byte) (MerkleNode, error){
	if h == nil || block == nil {
		return MerkleNode{}, nil
	}
	defer h.Reset()
	_, err := h.Write(block)
	if err != nil {
		return MerkleNode{}, err
	}
	return MerkleNode{
		HashValue: 		h.Sum(nil),
	}, nil
}

// 创建一棵树
func NewMerkleTree() *MerkleTree {
	return &MerkleTree{
		Size: 		0,
		Nodes: 		nil,
		Level: 		nil,
	}
}

//// 返回树的根
//func (merkleTree *MerkleTree) GetMerkleTreeRoot() *MerkleNode {
//	if merkleTree.Nodes == nil {
//		return nil
//	}
//	return merkleTree.Root
//}

// 返回树的高度
func (merkleTree *MerkleTree) GetMerkleTreeHeight() uint64 {
	return uint64(len(merkleTree.Level))
}

// 获取第 n 层的节点
func (merkleTree *MerkleTree) GetMerkleTreeNLevelNode(n uint64) []MerkleNode {
	if merkleTree.Level == nil || n == 0 || n > uint64(len(merkleTree.Level)) {
		return nil
	}
	return merkleTree.Level[n - 1]
}

// 返回 root
func (merkleTree *MerkleTree) GetMerkleTreeRoot() *MerkleNode {
	if merkleTree.Nodes == nil {
		return nil
	} else {
		return &merkleTree.Level[0][0]
	}
}

// 返回叶子节点
func (merkleTree *MerkleTree) GetMerkleTreeLeaves() []MerkleNode {
	if merkleTree.Level == nil {
		return nil
	}
	return merkleTree.Level[len(merkleTree.Level) - 1]
}

// 根据节点数量判断树的高度
func getMerkleTreeHeight(nodeCount uint64) uint64 {
	if nodeCount == 0 {
		return 0
	} else if nodeCount == 1 {
		return 2
	} else {
		return logBaseTwo(nextPowerOfTwo(nodeCount)) + 1
	}
}

// 计算节点数
func getNodeCount(height, size uint64) uint64 {
	if isPowerOfTwo(size) {
		return 2 * size - 1
	}
	count := size
	prev  := size
	for i := uint64(1); i < height; i++ {
		next := (prev + prev % 2) / 2
		count += next
		prev = next
	}
	return count
}

// 生成 MerkleTree
func (merkleTree *MerkleTree) Init(blocks [][]byte, h hash.Hash) error {
	blockCount := uint64(len(blocks))
	if blockCount == 0 {
		return errors.New("Empty tree")
	}
	height := getMerkleTreeHeight(blockCount)
	nodeCount := getNodeCount(height, blockCount)
	levels := make([][]MerkleNode, height)
	nodes  := make([]MerkleNode, nodeCount)
	size   := nodeCount

	// 创建叶子节点
	for k, v := range blocks {
		node, err := NewMerkleNode(h, v)
		if err != nil {
			return err
		}
		nodes[k] = node
	}
	levels[height - 1] = nodes[:len(blocks)]

	// 创建每个节点层
	current := nodes[len(blocks):]
	var i uint64
	for i = height - 1; i > 0; i-- {
		below := levels[i]
		wrote, err := merkleTree.initNodeLevel(below, current, h)
		if err != nil {
			return err
		}
		levels[i-1] = current[:wrote]
		current = current[wrote:]
	}
	merkleTree.Nodes = nodes
	merkleTree.Level = levels
	merkleTree.Size  = size

	return nil
}

// 生成所有非叶子节点
func (merkleTree *MerkleTree)initNodeLevel(below []MerkleNode, current []MerkleNode, h hash.Hash) (uint64, error) {
	h.Reset()
	size := h.Size()
	data := make([]byte, size*2)
	end  := (len(below) + len(below) % 2) / 2
	for i := 0; i < end; i++ {
		node := MerkleNode{}
		ileft  := 2 * i
		iright := 2 * i + 1
		left := &below[ileft]
		var right *MerkleNode = nil
		if len(below) > iright {
			right = &below[iright]
		}
		if right == nil {
			b := data[:size]
			copy(b, left.HashValue)
			node = MerkleNode{
				HashValue: 	b,
			}
		} else {
			copy(data[:size], below[ileft].HashValue)
			copy(data[size:], below[iright].HashValue)
			var err error
			node, err = NewMerkleNode(h, data)
			if err != nil {
				return 0, err
			}
		}
		// Point the new node to its children and save
		node.lchild = left
		node.rchild = right
		current[i] = node

		// Reset the data slice
		data = data[:]
	}
	return uint64(end), nil
}