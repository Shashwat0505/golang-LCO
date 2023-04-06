package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
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
	fmt.Println("Press  2 for deposit money-->2")
	fmt.Println("Press  3 for check balance-->3")
	fmt.Println("Press  4 for exit from here-->4")
	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		withdrawMoney(&currentuser)
		goto label1

	case 2:
		depositMoney(&currentuser)
		goto label1

	case 3:
		checkBalance(&currentuser)
		goto label1

	case 4:
		fmt.Println("!***Thank you for visiting***!")
		return
	default:
		fmt.Println("!---Please enter valid input---!")
		goto label1

	}

}

// --------------------------------withdrawMoney-------------------------------------------------------------------------------------------------
func withdrawMoney(u *user) {
	fmt.Println("Your current balance is ₹", u.balance)
label6:

	fmt.Println("!---Decimal values shold not be entered---!")
	fmt.Println("Enter the amount in multiple of 500 to withdraw:-")
	var withdrawAmont string
	fmt.Scan(&withdrawAmont)
	withdrawAmont = strings.ReplaceAll(withdrawAmont, ",", "")
	if !IsValid(withdrawAmont) {
		fmt.Println("Invalid amount!!Please try again")

	label7:
		fmt.Println("Do you want to continue??")
		fmt.Println("Press y for continue:-")
		fmt.Println("Press n for exit:-")
		var input1 string
		fmt.Scan(&input1)
		input1 = strings.ToLower(input1)
		switch input1 {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***")
			time.Sleep(time.Millisecond)
			os.Exit(0)

		default:
			fmt.Println("Invalid input")
			goto label7
		}

	}
	if checkforDot(withdrawAmont) {
		fmt.Println("Invalid amount!!Please try again")
		goto label6

	}
	tempBalance, _ := strconv.ParseFloat(u.balance, 64)
	withdrawcash, _ := strconv.ParseFloat(withdrawAmont, 64)

	if !IsValidWithdrawAmount(withdrawcash, tempBalance) {

		fmt.Println("!---Insufficient Balance---!")

	label8:
		fmt.Println("Do you want to continue??")

		fmt.Println("Press y for continue:-")
		fmt.Println("Press n for exit:-")
		var input2 string
		fmt.Scan(&input2)
		input2 = strings.ToLower(input2)
		switch input2 {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***")
			time.Sleep(time.Millisecond)
			os.Exit(0)

		default:
			fmt.Println("Invalid input")
			goto label8
		}
	} else {

		tempBalance = tempBalance - withdrawcash
		u.balance = fmt.Sprintf("%v", tempBalance)
		fmt.Println("Now your current balance is:-", u.balance)

	label9:
		fmt.Println("Do you want to continue?")
		fmt.Println("Press y for yes")
		fmt.Println("Press n for no")

		var choose string
		fmt.Scan(&choose)
		choose = strings.ToLower(choose)

		switch choose {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***!")
			time.Sleep(time.Millisecond)

			os.Exit(0)

		default:
			fmt.Println("Incorrect input!Please try again")
			goto label9

		}
	}
}

// --------------------------------depositMoney-------------------------------------------------------------------------------------------------
func depositMoney(u *user) {
	fmt.Println(u.UserName)
	fmt.Println("Your current balance is ₹", u.balance)
label3:
	fmt.Println("Enter the amount you want to deposit:-")
	var input string
	fmt.Scan(&input)
	input = strings.ReplaceAll(input, ",", "")

	if !IsValid(input) {
		fmt.Println("Invalid amount!!Please try again")
		goto label3

	}

	tempBalance, _ := strconv.ParseFloat(u.balance, 64)

	depositAmount, _ := strconv.ParseFloat(input, 64)

	if depositAmount > 100000 {
		fmt.Println("You can't enter this large amount", depositAmount)
		fmt.Println("!--Please reduce this amount---!")

		fmt.Println("Do you want to continue?")
		fmt.Println("Press y for yes")
		fmt.Println("Press n for no")

		var choose string
		fmt.Scan(&choose)
		choose = strings.ToLower(choose)
	label10:
		switch choose {
		case "y":
			return

		case "n":
			fmt.Println("!***Thank you for visiting***!")
			time.Sleep(time.Millisecond)

			os.Exit(0)

		default:
			fmt.Println("Incorrect input!Please try again")
			goto label10

		}

	}
	tempBalance += depositAmount

	u.balance = fmt.Sprintf("%v", tempBalance)

	fmt.Println("Now your balance is:-", u.balance)
label2:
	fmt.Println("Do you want to continue?")
	fmt.Println("Press y for yes")
	fmt.Println("Press n for no")

	var choose string
	fmt.Scan(&choose)
	choose = strings.ToLower(choose)

	switch choose {
	case "y":
		return

	case "n":
		fmt.Println("!***Thank you for visiting***!")
		time.Sleep(time.Millisecond)

		os.Exit(0)

	default:
		fmt.Println("Incorrect input!Please try again")
		goto label2

	}

}

// --------------------------------checkBalance-------------------------------------------------------------------------------------------------
func checkBalance(u *user) {
	fmt.Println("Your current account balance is:-", u.balance)
	fmt.Println("Do you want to continue?")
	fmt.Println("Press y for yes")
	fmt.Println("Press n for no")

	var choose string
	fmt.Scan(&choose)
	choose = strings.ToLower(choose)
label5:
	switch choose {
	case "y":
		return

	case "n":
		fmt.Println("!***Thank you for visiting***!")
		time.Sleep(time.Millisecond)

		os.Exit(0)

	default:
		fmt.Println("Incorrect input!Please try again")
		goto label5

	}

}

// --------------------------------IsValid-------------------------------------------------------------------------------------------------
func IsValid(n string) bool {

	for _, val := range n {
		if unicode.IsLetter(val) || unicode.IsSymbol(val) || unicode.IsSpace(val) || unicode.IsPunct(val) {
			return false
		}
	}
	return true
}

// --------------------------------CreateNewAccount-------------------------------------------------------------------------------------------------
func CreateNewAccount() {
	fmt.Println("Please enter username")
	var username string
	fmt.Scan(&username)
	fmt.Println("Please enter account number")
	var accountNo string
	fmt.Scan(&accountNo)
label12:
	fmt.Println("Please enter password for new Account")
	var password string
	fmt.Scan(&password)
	if !IsValidPin(password) {
		fmt.Println("!---PIN should be only digits and its length should be 4---!")
		goto label12

	}

label4:
	fmt.Println("Please enter amount of balance you want to add to account")
	var inputAmount string
	fmt.Scan(&inputAmount)
	inputAmount = strings.ReplaceAll(inputAmount, ",", "")
	if !IsValid(inputAmount) {
		fmt.Println("Invalid amount entered")
		goto label4
	}

	users = append(users, user{UserName: username, PIN: password, AccountNo: accountNo, balance: inputAmount})

	main()

}

// --------------------------------LoginToAccount-------------------------------------------------------------------------------------------------
func LogInToAccount() (user, error) {

	fmt.Println("Please enter accountno:-")
	var acno string
	fmt.Scan(&acno)

	fmt.Println("Please enter pin:-")
	var pin string
	fmt.Scan(&pin)

	for _, val := range users {
		if val.AccountNo == acno && val.PIN == pin {

			fmt.Println("!***Welcome ", val.UserName, "***!")

			return val, nil
		}
	}

	fmt.Println("Sorry invalid account-number or pin number")
	temp := user{}
	return temp, errors.New("Invalid credentials")

}

// --------------------------------checkfordots-------------------------------------------------------------------------------------------------
func checkforDot(s string) bool {
	if strings.Contains(s, ".") {
		return true
	}
	return false
}

// --------------------------------IsValidWithdrawAmount-------------------------------------------------------------------------------------------------

func IsValidWithdrawAmount(withdrawCash float64, balanceCash float64) bool {

	if balanceCash < withdrawCash {
		return false

	}

	withdrawtemp := int(withdrawCash)
	if withdrawtemp%500 != 0 {
		return false

	}
	if withdrawCash > 20000 {
		return false
	}
	return true

}

// --------------------------------IsValidPin-------------------------------------------------------------------------------------------------

func IsValidPin(s string) bool {
	if len(s) != 4 {
		return false
	}
	for _, val := range s {
		if unicode.IsLetter(val) || unicode.IsSymbol(val) || unicode.IsSpace(val) || unicode.IsPunct(val) {
			return false
		}
	}
	return true
}
