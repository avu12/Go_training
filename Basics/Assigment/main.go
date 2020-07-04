package main

import "fmt"

func main() {

	slice := []int{}
	for i := 0; i < 11; i++ {
		slice = append(slice, i)
	}
	for _, numb := range slice {
		if slice[numb]%2 == 0 {
			fmt.Println(slice[numb], "is even")
		} else {
			fmt.Println(slice[numb], "is odd")
		}
	}
}
