package onedoubleelementinarray

import (
	"fmt"
	"math/rand"
)

func CalculateDoubleElement(size int) {
	arr := make([]int, size)
	arrsum := 0
	nodoublesum := int((size - 2) * (size - 1) / 2)

	for i := 0; i < size-1; i++ {
		arr[i] = i
	}
	arr[size-1] = rand.Intn(size - 1)
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	for i := range arr {
		arrsum += arr[i]
	}

	fmt.Println(arrsum - nodoublesum)

}
