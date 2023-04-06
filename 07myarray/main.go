package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golangs")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println("fruit list is ", fruitList)

	var vegList = [3]string{"Potato", "Cucumber"}
	fmt.Println(vegList)

}
