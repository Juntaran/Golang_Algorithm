/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/13 18:00
  */

package main

import (
	"Golang_Algorithm/Data_Structure/Link_List"
	"fmt"
	"log"
)

func main() {
	testList := Link_List.New()
	testListB := Link_List.New()
	testList.PrintList()
	fmt.Println("Length is", testList.LenList())
	fmt.Println("isEmpty", testList.IsEmpty())

	testList.Prepend(1)
	testList.PrintList()
	fmt.Println("Length is", testList.LenList())
	fmt.Println("isEmpty", testList.IsEmpty())

	for i := 1; i <= 9; i++ {
		testList.Append(2)
	}
	testList.PrintList()
	fmt.Println("Length is", testList.LenList())
	fmt.Println("isEmpty", testList.IsEmpty())

	err := testList.Insert(3, 3)
	if err != nil {
		log.Println(err)
	}
	err = testList.Insert(0, 0)
	if err != nil {
		log.Println(err)
	}
	err = testList.Insert(7, 1)
	if err != nil {
		log.Println(err)
	}
	testList.PrintList()
	fmt.Println("Length is", testList.LenList())
	fmt.Println("isEmpty", testList.IsEmpty())

	err = testList.DeleteIndex(4)
	if err != nil {
		log.Println(err)
	}
	err = testList.DeleteIndex(16)
	if err != nil {
		log.Println(err)
	}
	testList.PrintList()
	fmt.Println("Length is", testList.LenList())
	fmt.Println("isEmpty", testList.IsEmpty())

	testList.DeleteValue(9)
	testList.PrintList()
	fmt.Println("Length is", testList.LenList())
	fmt.Println("isEmpty", testList.IsEmpty())

	testList.DeleteValue(7)
	testList.PrintList()
	fmt.Println("Length is", testList.LenList())
	fmt.Println("isEmpty", testList.IsEmpty())

	testList.DeleteValue(2)
	testList.PrintList()
	fmt.Println("Length is", testList.LenList())
	fmt.Println("isEmpty", testList.IsEmpty())

	testList.Clear()
	testList.PrintList()
	fmt.Println("Length is", testList.LenList())
	fmt.Println("isEmpty", testList.IsEmpty())

	testList.Connect(testListB)
	testList.PrintList()
	fmt.Println("Length is", testList.LenList())
	fmt.Println("isEmpty", testList.IsEmpty())

	for i := 1; i <= 9; i++ {
		testList.Append(2)
		testListB.Append(i)
	}
	testList.Connect(testListB)
	testList.PrintList()
	fmt.Println("Length is", testList.LenList())
	fmt.Println("isEmpty", testList.IsEmpty())

	element, err2 := testList.GetIndex(2)
	if err2 != nil {
		log.Println(err2)
	} else {
		fmt.Println(element)
	}

	element, err2 = testList.GetIndex(20)
	if err2 != nil {
		log.Println(err2)
	} else {
		fmt.Println(element)
	}

	ret, count := testList.FindIndex(2)
	fmt.Println(ret)
	fmt.Println(count)

	ret, count = testList.FindIndex(12)
	fmt.Println(ret)
	fmt.Println(count)

	err3 := testList.ChangeList(0, 5)
	if err3 != nil {
		log.Println(err3)
	} else {
		testList.PrintList()
		fmt.Println("Length is", testList.LenList())
		fmt.Println("isEmpty", testList.IsEmpty())
	}

	err3 = testList.ChangeList(5, 5)
	if err3 != nil {
		log.Println(err3)
	} else {
		testList.PrintList()
		fmt.Println("Length is", testList.LenList())
		fmt.Println("isEmpty", testList.IsEmpty())
	}
}