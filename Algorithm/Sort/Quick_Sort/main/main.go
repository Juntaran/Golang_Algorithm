/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/18 17:18
  */

package main

import (
	"fmt"

	"github.com/Juntaran/Golang_Algorithm/Algorithm/Sort/Quick_Sort"
)

func main() {
	testArray := []int{4, 2, 6, 7, 9, 5, 1, 3}
	Quick_Sort.Quick_Sort_Recurse_Single(testArray)
	//Quick_Sort.Quick_Sort_Recurse_Double(testArray)
	//Quick_Sort.Quick_Sort_Unrecurse(testArray)
	fmt.Println(testArray)
}
