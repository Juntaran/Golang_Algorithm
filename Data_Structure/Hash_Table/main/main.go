/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/14 06:47
  */

package main

import (
	"Golang_Algorithm/Data_Structure/Hash_Table"
	"fmt"
	"log"
)

func main() {
	hashTable := Hash_Table.New(5)
	hashTable.Insert("foo", "bar")
	hashTable.Insert("fiz", "buzz")

	fmt.Println("Size:", hashTable.ReturnSize())
	ret := hashTable.ContainsKey("foo1")
	fmt.Println(ret)

	val, err := hashTable.GetValue("foo6")
	if err == nil {
		fmt.Println(val)
	} else {
		log.Println(err)
	}

	hashTable.Insert("foo", "bomb")
	fmt.Println("Size:", hashTable.ReturnSize())
	val, err = hashTable.GetValue("foo")
	if err == nil {
		fmt.Println(val)
	} else {
		log.Println(err)
	}

	err = hashTable.Delete("foo")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Delete Size:", hashTable.ReturnSize())
	val, err = hashTable.GetValue("foo")
	if err == nil {
		fmt.Println(val)
	} else {
		log.Println(err)
	}
}