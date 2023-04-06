package main

import (
	"fmt"
	"strconv"
)

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	var ans string
	for _, val := range colors {
		switch val {
		case "Black":
			ans += "0"
		case "Brown":
			ans += "1"
		case "Red":
			ans += "2"
		case "Orange":
			ans += "3"
		case "Yellow":
			ans += "4"
		case "Green":
			ans += "5"
		case "Blue":
			ans += "6"
		case "Violet":
			ans += "7"
		case "Grey":
			ans += "8"
		case "White":
			ans += "9"

		}

	}

	temp, _ := strconv.Atoi(ans)
	return temp
	panic("Implement the Value function")
}

func main() {

	colors := []string{
		"Green",
		"Black",
	}
	t := Value(colors)
	fmt.Println(t)
}
