package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	doSomething(ctx)

}

func doSomething(ctx context.Context) {
	fmt.Println("here I am doing something")
}
