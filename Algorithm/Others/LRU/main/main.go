/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/12/9 21:23
  */

package main

import (
	"fmt"

	"github.com/Juntaran/Golang_Algorithm/Algorithm/Others/LRU"
)

func main() {
	lru := LRU.NewLRUCache(3)

	lru.Set(10,"value1")
	lru.Set(20,"value2")
	lru.Set(30,"value3")
	lru.Set(10,"value4")
	lru.Set(50,"value5")

	fmt.Println("LRU Size:",lru.Size())
	v, get, _ := lru.Get(30)
	if get {
		fmt.Println("Get(30) : ",v)
	}

	if lru.Remove(30) {
		fmt.Println("Remove(30) : true ")
	} else {
		fmt.Println("Remove(30) : false ")
	}
	fmt.Println("LRU Size:", lru.Size())
}
