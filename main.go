package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://x-team.com",
		"http://github.com",
		"http://stackoverflow.com",
		"http://google.com",
		"http://nonexistingurl.net",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkStatus(url, c)
	}

	for  u := range c{
		
		go func(url string){
			time.Sleep(time.Second * 3)
			checkStatus(url, c)
		}(u)

	}

}

func checkStatus(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is offline")
		c <- url
		return
	}
	fmt.Println(url, "is online")
	c <- url
}
