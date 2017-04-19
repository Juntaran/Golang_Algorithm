/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/19 10:54
  */

package main

import (
	"fmt"
	"Golang_Algorithm/Algorithm/Sort/Merge_Sort"
)

func main() {
	testArray := []int{4, 2, 6, 7, 9, 5, 1, 3}
	Merge_Sort.Merge_Sort(testArray)
	fmt.Println(testArray)
}
