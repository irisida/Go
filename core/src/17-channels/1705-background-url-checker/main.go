package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkURL(url string, c chan string) {
	res, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is down currently", url)
		s += fmt.Sprintf("Error: %v", err)
		fmt.Println(s)
		c <- url // sending url into the channel
	} else {
		defer res.Body.Close()

		s := fmt.Sprintf("%s -> status-code: %d\n", url, res.StatusCode)

		if res.StatusCode == 200 {
			s += fmt.Sprintf("%s checks OK", url)
			fmt.Println(s)
			c <- url
		}
	}
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://stackoverflow.com",
		"https://google.co.uk",
		"https://twitter.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkURL(url, c)
	}

	// syntax for channel iteration option 1
	// for {
	// 	go checkURL(<-c, c)
	// 	fmt.Println(strings.Repeat("-", 20))
	// 	time.Sleep(time.Second * 3)
	// }

	// syntax for channel iteration option 2
	// for url := range c {
	// 	time.Sleep(time.Second * 2)
	// 	go checkURL(url, c)
	// }

	// syntax for channel iteration option 3
	for url := range c {
		go func(u string) {
			time.Sleep(time.Second * 2)
			checkURL(u, c)
		}(url)
	}
}
