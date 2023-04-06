package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var ptr *int
	fmt.Println("Value of pointer is ", ptr)

	myNum := 23
	var ptr1 = &myNum
	fmt.Println("value of ptr1 is", ptr1)
	fmt.Println("value to which actual pointer is refering to", *ptr1)
	*ptr1 = *ptr1 + 1
	fmt.Println("now the value is ", myNum)
	fmt.Println("value of ptr1 is still", ptr1)
}
