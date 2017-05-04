/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/5 01:01
  */

package main

import (
	"Golang_Algorithm/Data_Structure/Set"
	"fmt"
	"log"
)

func main() {
	testSet := Set.New()
	testSet.PrintSet()
	fmt.Println("Length is", testSet.LenSet())
	fmt.Println("isEmpty:", testSet.IsEmpty())

	for i := 1; i <= 10; i++ {
		testSet.Add(i)
	}
	testSet.PrintSet()
	fmt.Println("Length is", testSet.LenSet())
	fmt.Println("isEmpty:", testSet.IsEmpty())

	fmt.Println(testSet.IsExist(3))
	fmt.Println(testSet.IsExist(11))

	err := testSet.Delete(3)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(testSet.IsExist(3))
	fmt.Println("Length is", testSet.LenSet())
	fmt.Println("isEmpty:", testSet.IsEmpty())

	err = testSet.Delete(3)
	if err != nil {
		log.Println(err)
	}
}
