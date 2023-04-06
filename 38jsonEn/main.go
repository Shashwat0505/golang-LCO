package main

import (
	"encoding/json"
	"fmt"
)

type Permissions map[string]bool

type user struct {
	Name        string `json:"username"`
	Password    string `json:"-"`
	Permissions `json:",omitempty"`
}

func main() {
	users := []user{
		{"insac", "1234", nil},
		{"god", "42", Permissions{"admin": true}},
		{"devil", "666", Permissions{"write": true}},
	}
	out, err := json.MarshalIndent(users, " ", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

}
