package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius int
}

type Square struct {
	length int
}

func (c Circle) Area() float64 {

	return float64(c.radius) * float64(c.radius) * 3.14

}

func main() {
	c1 := Circle{5}

	fmt.Println(c1.Area())

}
