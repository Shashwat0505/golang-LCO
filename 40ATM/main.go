package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
	"40ATM/functions"
)

type user struct {
	UserName  string
	PIN       string
	AccountNo string
	balance   string
}

var users = []user{
	user{UserName: "Shashwat", PIN: "5522", AccountNo: "555", balance: "5500"},
}

var balanceAmount float64

// --------------------------------main-------------------------------------------------------------------------------------------------
func main() {
	var currentuser user
	fmt.Println("!***Welcome to ATM***!")
	fmt.Println(strings.Repeat("*", 22))
label0:
	fmt.Println("Already have an account or create new one??")
	fmt.Println("Press n for new account OR Press l for login to account")
	var log string
	fmt.Scan(&log)
	log = strings.TrimSpace(log)
	log = strings.ToLower(log)

	switch log {
	case "n":
		CreateNewAccount()

	case "l":
	loginLabel:
		cu, err := LogInToAccount()
		if err != nil {
			fmt.Println("Invalid credentials")
			goto loginLabel

		}
		currentuser = cu

	default:
		fmt.Println("Invalid input try again")
		goto label0
	}
label1:
	fmt.Println("Press  1 for withdraw money-->1")
	fmt.Println(strings.Repeat("-", 31))
	fmt.Println("Press  2 for deposit money-->2")
	fmt.Println(strings.Repeat("-", 31))
	fmt.Println("Press  3 for check balance-->3")
	fmt.Println(strings.Repeat("-", 31))
	fmt.Println("Press  4 for exit from here-->4")
	fmt.Println(strings.Repeat("-", 31))
	var input string
	fmt.Scan(&input)

	switch input {
	case "1":
		WithdrawMoney(&currentuser)
		goto label1

	case "2":
		DepositMoney(&currentuser)
		goto label1

	case "3":
		CheckBalance(&currentuser)
		goto label1

	case "4":
		fmt.Println("!***Thank you for visiting***!")
		time.Sleep(time.Millisecond)
		os.Exit(0)

	default:
		fmt.Println("!---Please enter valid input---!")
		goto label1

	}

}

// --------------------------------withdrawMoney-------------------------------------------------------------------------------------------------

// --------------------------------DepositMoney-------------------------------------------------------------------------------------------------

// --------------------------------CheckBalance-------------------------------------------------------------------------------------------------

// --------------------------------IsValid-------------------------------------------------------------------------------------------------
func IsValid(n string) bool {

	for _, val := range n {
		if unicode.IsLetter(val) || unicode.IsSymbol(val) || unicode.IsSpace(val) || unicode.IsPunct(val) || unicode.IsPunct(val) {
			return false
		}
	}
	return true
}

// --------------------------------CreateNewAccount-------------------------------------------------------------------------------------------------

// --------------------------------LoginToAccount-------------------------------------------------------------------------------------------------

// --------------------------------CheckforDots-------------------------------------------------------------------------------------------------

// --------------------------------IsValidWithdrawAmount-------------------------------------------------------------------------------------------------

// --------------------------------IsValidPin-------------------------------------------------------------------------------------------------
