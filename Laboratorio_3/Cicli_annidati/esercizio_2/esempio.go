package main

import "fmt"

func main() {
	n := 3

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for z := 0; z < n; z++ {
				fmt.Print("*")
			}
		}
	}
	fmt.Println()
}
