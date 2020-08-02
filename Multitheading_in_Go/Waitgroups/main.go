package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches   []string
	waitgroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
)

func fileSearch(root string, filename string) {
	fmt.Println("Searching in", root)
	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println(err)

	}
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			waitgroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}

	}
	waitgroup.Done()

}

func main() {
	waitgroup.Add(1)
	go fileSearch("D:/", "main.go")
	waitgroup.Wait()
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
}
