package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.com",
		"http://stackoverflow.com",
	}
	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel) //BUG: if main routine finish early, exits automatically Need to use channel!
	}
	for l := range channel {
		go func(link string) { //function literal !
			time.Sleep(2 * time.Second)
			checkLink(link, channel)
		}(l)

	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "DOWN!")
		c <- link
		return
	}
	fmt.Println(link + " is UP!")
	c <- link
}
