package main

import "fmt"

func main() {
	fmt.Println("Welcome to the defer")

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

}
