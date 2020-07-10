package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("santa_helper.txt")
	lvl := 0
	if err != nil {
		log.Fatal(err)
	}
	for i := range content {
		if content[i] == 40 {
			lvl++
		} else if content[i] == 41 {
			lvl--
		} else {
			os.Exit(-2)
		}
	}

	fmt.Println(lvl)
	//fmt.Printf("File contents: %s", content)
}
