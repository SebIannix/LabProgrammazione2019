package main

import "fmt"

func main() {
	n := 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ { // for j := 0; j < n; j-- { : il for precedente andava in loop infinito dato che
			// la variabile j veniva decrementata invece di essere incrementata.
			// Quindi, la condizione j<n era sempre vera
			fmt.Print("*")
		}
	}
	fmt.Println()
}
