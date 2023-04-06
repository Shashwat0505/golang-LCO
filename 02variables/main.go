package main

import "fmt"

func main() {
	var isLoggedIn bool=false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	var smallval uint8=255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type:%T\n",smallval)

	var temp string
	fmt.Println(temp)
	fmt.Printf("type of string is:%T\n",temp)
	
	var num int
	fmt.Println(num)
	fmt.Printf("type of num is:%T\n",num)

	 token:=2000
	fmt.Println(token)
}