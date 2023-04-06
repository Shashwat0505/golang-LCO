package main

import (
	"fmt"
	"math/rand"
)

func main() {

	arr := [5]string{"Sad", "Angry", "Happy", "Amazing", "Terrible"}
	fmt.Println("Please enter your name:-")
	var name string
	fmt.Scan(&name)

	fmt.Println(name, "is ", arr[rand.Intn(len(arr))])

}
