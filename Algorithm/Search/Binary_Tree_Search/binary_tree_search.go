/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/25 14:13
  */

package Binary_Tree_Search

import (
	"Golang_Algorithm/Data_Structure/Binary_Tree"
	"fmt"
)

/*
	递归查找二叉排序树T中是否存在key,指针f指向T的双亲，其初始调用值为NULL
	若查找成功，则指针p指向该数据元素结点，并返回1
	否则指针p指向查找路径上访问的最后一个结点并返回0
*/

func Binary_Tree_Search(nums []int, key int) bool {
	bst := Binary_Tree.NewTree(nil)
	for i := 0; i < len(nums); i++ {
		bst.Insert(nums[i])
	}
	//bst.LevelTraverse()
	_, err := bst.Search(key)
	if err != nil {
		fmt.Println("Not Find")
		return false
	} else {
		fmt.Println("Find key")
		return true
	}
}
