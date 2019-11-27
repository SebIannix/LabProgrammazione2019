package main

import "fmt"

func main() {
	var a = [4]int{1, 2, 3, 4}
	var b = [...]rune{'a', 'b', 'c'}

	fmt.Println(a)

	for i:=0; i<len(a); i++ {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}

	fmt.Println()

	for i, v := range b {
		fmt.Println("Indice", i, " - Valore:", v)
	}

}