package main

import "fmt"

func main() {
	var color string
	fmt.Scan(&color)
	value := colorCode(color)
	fmt.Println(value)

}

func colors() []string {
	colorsList := []string{
		"Black",
		"Brown",
		"Red",
		"Orange",
		"Yellow",
		"Green",
		"Blue",
		"Violet",
		"Grey",
		"White",
	}
	return colorsList
}

func colorCode(color string) int {
	cl := colors()
	for i := 0; i < len(cl); i++ {
		if cl[i] == color {
			return i
		}
	}
	return 0

}
