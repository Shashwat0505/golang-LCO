package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in go")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["Py"] = "python"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS shorts for:", languages["JS"])
	delete(languages, "Py")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("for key %v,the value is %v\n", key, value)
	}

}
