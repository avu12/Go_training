package scheluder

import (
	"fmt"
	"sync"
	"time"
)

// You have 3 thread, each write 111111... 222222.... and 333333.... continously, write a program, with the thread output 123123123.....
var (
	waitgroup     = sync.WaitGroup{}
	lock          = sync.Mutex{}
	condvar       = sync.NewCond(&lock)
	lastused  int = 3
)

func Scheluder() {

	go PrintOne()
	go PrintTwo()
	go PrintThree()

}

func PrintOne() {
	for {
		lock.Lock()
		for lastused != 3 {
			condvar.Wait()
		}
		fmt.Print(1)
		lastused = 1
		condvar.Broadcast()
		lock.Unlock()
		time.Sleep(time.Second / 100)
	}

}
func PrintTwo() {
	for {
		lock.Lock()
		for lastused != 1 {
			condvar.Wait()
		}
		fmt.Print(2)
		lastused = 2
		condvar.Broadcast()
		lock.Unlock()
		time.Sleep(time.Second / 100)
	}
}
func PrintThree() {
	for {
		lock.Lock()
		for lastused != 2 {
			condvar.Wait()
		}
		fmt.Print(3)
		lastused = 3
		condvar.Broadcast()
		lock.Unlock()
		time.Sleep(time.Second / 100)
	}
}
