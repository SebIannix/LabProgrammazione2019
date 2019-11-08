package main

import "fmt"

func main() {
	var a int = 4
	var b float64 = 12.5
	var e int // è necessario specificare un tipo per la variabile e
	// var c int : possiamo rimuovere la dichiarazione della variabile c
	c := float64(a) + b // le operazioni devono essere fatte tra variabili dello stesso tipo
	d := float64(a) / b
	fmt.Println(a, b, c, d, e) // la variabile f non esiste e va sostituita con la variabile e
}
