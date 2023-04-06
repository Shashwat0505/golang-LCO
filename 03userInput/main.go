package main

import (
	"fmt"
	"math/rand"
	"time"
)

type user struct {
	UserName string
	PIN      string
	balance  string
}

var users = []user{
	user{UserName: "Shashwat", PIN: "5522", balance: "5500"},
}

func main() {

	// welcome := "Welcome to user input"
	// fmt.Println(welcome)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter the rating for service:")

	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thanks for rating,", input)
	// fmt.Printf("the type of input is %T", input)

	// fmt.Println("Please enter username")
	// var username string
	// fmt.Scan(&username)
	// fmt.Println("Please enter password for new Account")
	// var password string
	// fmt.Scan(&password)

	// fmt.Println("Please enter amount of balance you want to add to account")
	// var inputAmount string
	// fmt.Scan(&inputAmount)

	// users = append(users, user{UserName: username, PIN: password, balance: inputAmount})
	// fmt.Println(users)

	min := 1000000000
	max := 9999999999
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(max-min) + min)

}
