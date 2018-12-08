/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 17:18
  */

package main

import (
	"Golang_Algorithm/Algorithm/Sort/Quick_Sort"
	"fmt"
)

func main() {
	testArray := []int{4, 2, 6, 7, 9, 5, 1, 3}
	Quick_Sort.Quick_Sort_Recurse(testArray)
	//Quick_Sort.Quick_Sort_Unrecurse(testArray)
	fmt.Println(testArray)
}
