package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")
	var fruitList = []string{"Mango", "Tomato", "Banana"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "pineapple", "chickoo")
	fmt.Println(fruitList)
	fruitList = append(fruitList[0:2])
	fmt.Println(fruitList)

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 236
	highscore[2] = 237

	fmt.Println(highscore)
	highscore = append(highscore, 1000, 444, 888)
	fmt.Println(highscore)
	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	var languages = []string{"Java", "Python", "Go", "C", "JavaScript"}
	fmt.Println(languages)
	var index int = 2
	languages = append(languages[:index], languages[index+1:]...)
	fmt.Println(languages)

}
