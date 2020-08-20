package moreduplicatedarray

import "fmt"

func Findaduplicate(arr []int) string {
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 || arr[i] > len(arr)-1 {
			return "Negative or too big numbers!"
		}
	}

	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == i {
			continue
		}
		if arr[i] == arr[arr[i]] {
			return fmt.Sprintf("Duplicate found %d", arr[i])
		}

		arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
		i--
	}
	return ("No duplicate in array!")
}
