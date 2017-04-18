/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/17 15:53
  */

package main

import (
	"Golang_Algorithm/Data_Structure/Binary_Tree"
	"fmt"
	"log"
)

func main() {
	n := Binary_Tree.NewNode(1)
	tree := Binary_Tree.NewTree(n)

	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(4)

	tree.PreOrderTraverse()
	tree.InOrderTraverse()
	tree.PostOrderTraverse()
	tree.LevelTraverse()

	fmt.Println("size:", tree.Size)

	m, err := tree.Search(2)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(m.Value)
	}
	err = tree.Delete(4)
	if err != nil {
		log.Println(err)
	}

	m, err = tree.Search(6)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(m.Value)
	}
	fmt.Println("size:", tree.Size)
}
