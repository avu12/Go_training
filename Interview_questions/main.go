package main

import (
	"fmt"
	"sync"

	"./palindrome"
	"./scheluder"
)

var (
	waitgroup = sync.WaitGroup{}
	//lock      = sync.Mutex{}
)

func main() {

	if false {
		fmt.Println(palindrome.IsPalindrome("racecar"))
	}
	waitgroup.Add(1)
	scheluder.Scheluder()
	waitgroup.Wait()

}
