package main

import (
	"sync"
	"time"
)

var (
	money = 100
	lock  = sync.Mutex{}
)

func add() {
	for i := 1; i < 1000000; i++ {
		lock.Lock()
		money += 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Add done")
}
func substract() {
	for i := 1; i < 1000000; i++ {
		lock.Lock()
		money -= 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Subtract done")
}

func main() {
	go add()
	go substract()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}
