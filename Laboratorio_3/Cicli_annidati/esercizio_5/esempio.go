package main

import "fmt"

func main() {
	n := 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j-- {
			fmt.Print("*")
		}
	}
	fmt.Println()
}
