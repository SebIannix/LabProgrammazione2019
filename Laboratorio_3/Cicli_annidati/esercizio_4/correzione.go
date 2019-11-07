package main

import "fmt"

func main() {
	n := 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ { // for j := 0; j < n; i++ { : il for precedente incrementava la variabile i
			// invece di incrementare il valore della variabile j, che quindi rimaneva sempre 0
			// La condizione j<n era quindi sempre vera
			fmt.Print("*")
		}
	}
	fmt.Println()
}
