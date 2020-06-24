package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		filename := os.Args[1]
		file, _ := os.Open(filename)
		bs := make([]byte, 999999)
		file.Read(bs)
		fmt.Println(string(bs))
		file.Close()
	}

}
