package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	t := time.Now()
	rand.Seed(t.UnixNano())
	var guess int = 10
	fmt.Scan(&guess)
	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d", n)
	}

}
