package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"

	"./mergesortedarrays"
	"./onedoubleelementinarray"
	"./palindrome"
	"./scheluder"
)

var (
	waitgroup = sync.WaitGroup{}
)

func main() {

	if false {
		fmt.Println(palindrome.IsPalindrome("racecar"))
	}
	if false {
		//for infinite run:
		waitgroup.Add(1)
		scheluder.Scheluder()
		waitgroup.Wait()
	}
	if false {
		arr1 := []int{}
		for i := 0; i < 1000000; i++ {
			arr1 = append(arr1, rand.Intn(500000))
		}
		arr2 := []int{}
		for i := 0; i < 1000000; i++ {
			arr2 = append(arr2, rand.Intn(5000000))
		}
		sort.Sort(sort.IntSlice(arr1))

		sort.Sort(sort.IntSlice(arr2))

		mergesortedarrays.MergesortedArrays(arr1, arr2)
	}

	onedoubleelementinarray.CalculateDoubleElement(90)

}
