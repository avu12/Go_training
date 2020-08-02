package main

import (
	"sync"
	"time"
)

var (
	money          = 100
	lock           = sync.Mutex{}
	moneyDeposited = sync.NewCond(&lock)
)

func add() {
	for i := 1; i < 1000; i++ {
		lock.Lock()
		money += 10
		println("Balance after addition:", money)
		moneyDeposited.Signal()
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Add done")
}
func substract() {
	for i := 1; i < 1000; i++ {
		lock.Lock()
		for money-20 < 0 {
			moneyDeposited.Wait()
		}
		money -= 20
		println("Balance after subtraction:", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Subtract done")
}

func main() {
	go add()
	go substract()
	time.Sleep(30000 * time.Millisecond)
	print(money)
}
