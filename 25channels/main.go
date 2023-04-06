package main

import "fmt"

func main() {
	fmt.Println("Welcome to channels")

	square := make(chan int)
	cube := make(chan int)

	num := 123
	go findsq(num, square)
	go findcube(num, cube)
	sqa, cuba := <-square, <-cube
	fmt.Println("value of square is", sqa)
	fmt.Println("value of cube is", cuba)

}
func findsq(number int, sq chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	sq <- sum
}

func findcube(number int, cub chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10

	}
	cub <- sum
}
