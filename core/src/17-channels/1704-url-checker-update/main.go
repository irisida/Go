package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func checkAndSaveBody(url string, dirName string, c chan string) {
	res, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is down currently", url)
		s += fmt.Sprintf("Error: %v", err)
		c <- s // sending into the channel
	} else {
		defer res.Body.Close()

		s := fmt.Sprintf("%s -> status-code: %d\n", url, res.StatusCode)

		if res.StatusCode == 200 {
			// read the response body as a byte slice that will be written to file
			bs, err := ioutil.ReadAll(res.Body)

			// create the file with the domain and a .txtsuffix
			file := dirName + "/" + strings.Split(url, "//")[1]
			file += ".txt"

			s += fmt.Sprintf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bs, 0664)
			if err != nil {
				s += "Error writing file"
				c <- s
			}
			s += fmt.Sprintf("%s checks OK", url)
			c <- s
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

	dirName := "responseBodies"

	// create a holding pen drectory for the responses
	// pass the name to the function called in the loop
	// so it can be appended to the full path for the
	// written output
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	c := make(chan string)

	for _, url := range urls {
		go checkAndSaveBody(url, dirName, c)
		fmt.Println(strings.Repeat("-", 20))
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}
