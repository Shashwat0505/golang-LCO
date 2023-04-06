package main

import (
	"errors"
	"fmt"
)

func FromRNA(rna string) ([]string, error) {

	ans := []string{}

	for i := 0; i < len(rna); {
		var temp string
		for j := i; j < i+3; j++ {

			temp += string(rna[j])
		}
		fmt.Println(temp)
		t1, _ := FromCodon(temp)
		if t1 == "STOP" {
			return ans, nil
		}
		ans = append(ans, t1)
		i = i + 3

	}
	return ans, nil

	panic("Please implement the FromRNA function")
}

func FromCodon(codon string) (string, error) {
	m1 := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAC": "Tyrosine",
		"UAU": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	if m1[codon] != "" {
		return m1[codon], nil
	}
	return m1[codon], errors.New("error")
	panic("Please implement the FromCodon function")
}

func main() {
	a1, err := FromRNA("UGGUGUUAUUAAUGGUUU")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a1)
	}
}
