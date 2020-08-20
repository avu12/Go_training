package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"

	"./firstduplicationarray"
	"./mergesortedarrays"
	"./moreduplicatedarray"
	"./onedoubleelementinarray"
	"./palindrome"
	"./scheluder"
)

func main() {
	arr := []int{2, 5, 5, 2, 3, 5, 1, 2, 4}
	firstduplicationarray.FindFirstDuplicate(arr)
}

func Exercises() {

	var (
		waitgroup = sync.WaitGroup{}
	)

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
	if false {
		onedoubleelementinarray.CalculateDoubleElement(90)
	}

	if false {
		arr := []int{2, 3, 1, 0, 2}
		fmt.Println(moreduplicatedarray.Findaduplicate(arr))
	}

}
