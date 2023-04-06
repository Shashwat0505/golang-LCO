package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var sum int = 0
var length int = 10
var j int = 1

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}

	if strings.HasSuffix(isbn, "X") {
		isbn = strings.ReplaceAll(isbn, "X", "10")
		sum = 10
		isbn = strings.TrimSuffix(isbn, "10")
		length = 9
		j = 2

	}

	for _, val := range isbn {
		if !unicode.IsNumber(val) {
			return false

		}
	}
	num, _ := strconv.Atoi(isbn)

	for i := 0; i < length; i++ {
		k := num % 10

		sum += k * j
		j++
		num = num / 10
	}
	if sum%11 == 0 {
		return true
	}
	return false

	panic("Please implement the IsValidISBN function")
}

func main() {
	t := IsValidISBN("3598215088")
	fmt.Println(t)

}
