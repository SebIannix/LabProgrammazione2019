package main

import (
	"fmt"
)

func main() {

	var stringaInput string

	fmt.Print("Parola in input: ")
	fmt.Scan(&stringaInput)

	for i := 0; i < (len(stringaInput)+1)/2; i++ {
		fmt.Println("Sottostringa", i+1, stringaInput[i:len(stringaInput)-i])
	}

	/*
		// Codice alternativo:
		for i := range stringaInput[:(len(stringaInput)+1)/2] {
			fmt.Println("Sottostringa", i+1, stringaInput[i:len(stringaInput)-i])
		}
	*/

}
