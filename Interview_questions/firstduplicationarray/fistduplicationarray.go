package firstduplicationarray

import "fmt"

func FindFirstDuplicate(arr []int) {

	hash := make(map[int]bool, 0)

	for _, v := range arr {
		if hash[v] == true {
			fmt.Printf("Found %d number second time!", v)
			return

		} else {
			hash[v] = true
		}

	}
	fmt.Println("No douplicate!")

}
