package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	weblist := []string{
		"https://google.com",

		"https://github.com",
	}

	for _, value := range weblist {
		go getStatusCode(value)
		wg.Add(1)

	}
	wg.Wait()

}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("oops in endpoint")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

}
