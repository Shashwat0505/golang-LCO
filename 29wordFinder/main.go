package main

import (
	"fmt"
	"os"
	"strings"
)

const db = "a quick brown fox jumps over a lazy dog fox dog"

func main() {

	word := strings.Fields(db)
	query := os.Args[1:]
	// var str string
	// fmt.Scanln(&str)
	// fmt.Println(str)
	// fmt.Println(word)
	// fmt.Println(query)
outer:
	for _, val := range query {
		for index, w := range word {
			if w == val {
				fmt.Printf("%v is found at index %d\n ", w, index)
				break outer
			}
		}
	}

}
