package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Enter the size of slice:-")
	var size int
	fmt.Scan(&size)
	fmt.Println("Please enter the range:-")
	var maxi int
	fmt.Scan(&maxi)
	var sl []int

loop:
	for i := 0; i < size; {
		n := rand.Intn(maxi)
		for _, val := range sl {
			if val == n {
				continue loop
			}
		}
		sl = append(sl, n)
		i++
	}
	fmt.Println(sl)

}
