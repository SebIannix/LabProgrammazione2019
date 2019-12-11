package main

import "fmt"

func main() {
	var a, b int = 10, 20
	var ptr1, ptr2 *int
	ptr1 = a
	*ptr1 += 10
	*ptr2 = 100
	ptr2 = &b
	fmt.Println(a, b)
}
