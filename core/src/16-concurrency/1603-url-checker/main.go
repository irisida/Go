package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func checkAndSaveBody(url string, dirName string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s is down at the minute\n", url)
		log.Println(err)
	} else {
		defer res.Body.Close()

		fmt.Printf("%s -> status-code: %d\n", url, res.StatusCode)

		if res.StatusCode == 200 {
			// read the response body as a byte slice that will be written to file
			bs, err := ioutil.ReadAll(res.Body)

			// create the file with the domain and a .txtsuffix
			file := dirName + "/" + strings.Split(url, "//")[1]
			file += ".txt"

			fmt.Printf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bs, 0664)
			if err != nil {
				log.Fatal(err)
			}
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

	for _, url := range urls {
		checkAndSaveBody(url, dirName)
		fmt.Println(strings.Repeat("-", 20))
	}
}
