/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/1/30 14:12
  */

package main


import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"Golang_Algorithm/Data_Structure/Merkle_Tree"
)

func splitData(data []byte, size int) [][]byte {
	// Splits data into an array of slices of len(size)
	count := len(data) / size
	blocks := make([][]byte, 0, count)
	for i := 0; i < count; i++ {
		block := data[i*size : (i+1)*size]
		blocks = append(blocks, block)
	}
	if len(data) % size != 0 {
		blocks = append(blocks, data[len(blocks)*size:])
	}
	return blocks
}

func main() {
	// Grab some data to make the tree out of, and partition
	data, err := ioutil.ReadFile("Data_Structure/Merkle_Tree/testdata.txt") // assume testdata.txt exists
	if err != nil {
		fmt.Println(err)
		return
	}
	blocks := splitData(data, 32)

	// Create & generate the tree
	tree := Merkle_Tree.NewMerkleTree()
	err = tree.Init(blocks, md5.New())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Height: %d\n", tree.GetMerkleTreeHeight())
	fmt.Printf("Root: %v\n", tree.GetMerkleTreeRoot())
	fmt.Printf("N Leaves: %v\n", len(tree.GetMerkleTreeLeaves()))
	fmt.Printf("Height 2: %v\n", tree.GetMerkleTreeNLevelNode(2))
	fmt.Println(tree.Level)
}