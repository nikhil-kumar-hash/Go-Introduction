package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}

	c := make(chan string) // creating a channel to pass string

	for _, link := range links {
		go checkLink(link, c)
	}

	// main routine is waiting for a msg to pop
	// in the channel , then as soon as we get one it will exit
	// Recieving msgs from a channel is a blocking call

	for l := range c { // infinite loop
		// blocking operation
		go func(t string) {
			time.Sleep(5 * time.Second)
			checkLink(t, c)
		}(l) // function literal

	}
}

func checkLink(link string, c chan string) bool {

	_, err := http.Get(link)
	if err != nil {
		//fmt.Println("Error :", err)
		fmt.Println(link + " is down")
		c <- link
		return false
	}
	println(link + "is up and running")
	c <- link
	return true
}

// sending data through a channel

// 1. channel <- 5  // --->  Send the value 5 into the channel
// 2. myNumber <- channel // -----> wait for a value to be sent into the channel.When we get the value,assign it to"myNumber variable
// 3. fmt.Println(<-channel) // wait for a value to be sent into the channel. When we get one, log it out immediately.
