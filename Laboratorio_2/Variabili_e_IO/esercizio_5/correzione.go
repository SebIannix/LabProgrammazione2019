package main

import "fmt"

func main() {
	var a int // var a : manca il tipo
	a = 10

	// var b int : variabile b dichiarata 2 volte
	b := 20

	c := 30 // c = 30 : variabile c non dichiarata

	d := a + b + c

	fmt.Println(d) // fmt.Println((d) : una parentesi di troppo
}
