package main

import "fmt"

func main() {
	fmt.Println("Welcome to fuctions")
	greeter()
	result := adder(3, 4)
	fmt.Println("result is:-", result)

	proResult, hi := proRes(4, 5, 6, 7, 1, 3, 4, 5, 4)
	fmt.Println("the pro result is:-", proResult)
	fmt.Println("the string is:-", hi)

}

func greeter() {
	fmt.Println("Namaste golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo

}

func proRes(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hello world"
}
