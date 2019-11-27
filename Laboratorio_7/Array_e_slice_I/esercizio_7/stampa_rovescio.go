package main

import (
	"fmt"
)

func main() {
	// leggo il primo intero n
	var n int
	fmt.Scan(&n)

	// leggo n interi e li memorizzo in una slice
	fmt.Printf("Inserisci %d numeri:\n", n)
	numeri := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&numeri[i])
	}

	fmt.Println("Numeri in ordine inverso:")
	for i := range numeri {
		fmt.Print(numeri[len(numeri)-1-i], " ")
	}
	fmt.Println()

}