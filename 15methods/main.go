package main

import "fmt"

func main() {
	shashwat := User{"Shashwat", 19, "sm@gmail.com", true}
	fmt.Println(shashwat)
	shashwat.GetStatus()
	fmt.Println(shashwat.Email)

}

type User struct {
	Name   string
	Age    int
	Email  string
	status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.status)
	u.Email = "new@mail.com"
	fmt.Println("new email for user is:-", u.Email)

}
