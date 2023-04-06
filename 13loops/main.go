package main

import "fmt"

func main() {
	fmt.Println("welcome to main")
	days := []string{"Monday", "Tuesday", "wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	for i := range days {
		fmt.Println(days[i])

	}

	for index, day := range days {
		fmt.Printf("the index is %v and value is %v\n", index, day)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 5 {
			goto lca

		}
		fmt.Println("value is :", rougueValue)

		rougueValue++
	}

lca:
	{
		fmt.Println("welcome to lca")
	}
}
