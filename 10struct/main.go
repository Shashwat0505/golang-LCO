package main

import "fmt"

func main() {
	fmt.Println("you are inside struct")

	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

	shashwat := User{"Shashwat", "shashwat@go.dev", true, 18}
	fmt.Println(shashwat)
	fmt.Printf("shashwat details are:-%+v\n", shashwat)
	fmt.Printf("email of shashwat is %v and status of shashwat is %v", shashwat.Email, shashwat.Status)

}
