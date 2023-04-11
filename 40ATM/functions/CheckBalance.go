package functions

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func CheckBalance(u *user) {
	fmt.Println("Your current account balance is:-₹", u.balance)
label5:
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
		goto label5

	}

}
