package mergesortedarrays

import "fmt"

func MergesortedArrays(arr1 []int, arr2 []int) {
	var merged []int
	var arr1index = 0
	var arr2index = 0

	for i := 0; i < len(arr1)+len(arr2); i++ {
		if arr1[arr1index] < arr2[arr2index] && arr1index < len(arr1)-1 {
			merged = append(merged, arr1[arr1index])
			arr1index++
			if arr1index == len(arr1)-1 {
				merged = append(merged, arr2...)
				break
			}
		} else if arr2index < len(arr2)-1 {
			merged = append(merged, arr2[arr2index])
			arr2index++
			if arr2index == len(arr2)-1 {
				merged = append(merged, arr1...)
				break
			}
		}

	}

	//fmt.Println("MERGED:", merged)
	fmt.Println("MERGED")

}
