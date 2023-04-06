package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex
var signals = []string{}

func main() {

	signals = append(signals, "test")
	sites := []string{
		"https://google.com",
		"https://github.com",
	}
	for _, values := range sites {
		go getStatus(values)
		wg.Add(1)

	}
	wg.Wait()
	fmt.Println(signals)

}

func getStatus(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("oops in endpoint")
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d is the statuscode and the endpoint is %s\n", res.StatusCode, endpoint)

}
