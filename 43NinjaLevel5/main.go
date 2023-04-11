package main

import (
	"fmt"
	"time"
)

func main() {

	go hello()
	go bye()

	time.Sleep(time.Second)
	fmt.Println("Hello aliens")

	fmt.Println("Enter the n for fibonacci:-")
	var n int
	fmt.Scan(&n)

	ch := make(chan int, n)
	go fibo(ch, n)

	for val := range ch {
		fmt.Println(val)
	}
}

func hello() {
	fmt.Println("Inside hello function")
}
func bye() {
	fmt.Println("inside bye function")

}
func fibo(ch chan int, n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
