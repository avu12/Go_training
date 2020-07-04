package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		filename := os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error")
			os.Exit(19)
		}
		bs := make([]byte, 999999)
		file.Read(bs)
		fmt.Println(string(bs))
		file.Close()
	}

}
