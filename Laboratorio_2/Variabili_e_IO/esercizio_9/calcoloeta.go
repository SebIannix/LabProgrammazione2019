package main

import "fmt"

func main() {
	var età1, età2 int

	// leggo da standard input le due età
	fmt.Print("Età persona 1 = ")
	fmt.Scanln(&età1)
	fmt.Print("Età persona 2 = ")
	fmt.Scanln(&età2)

	somma := età1 + età2

	// per calcolare correttamente la media è necessario convertire la somma delle età da int a float64:
	// senza questa conversione, si otterrebbe il risultato della divisione tra interi, perdendo la parte decimale
	media := float64(età1+età2) / 2

	fmt.Println("Somma età =", somma)
	fmt.Println("Media età =", media)

	età1 += 10
	età2 += 10

	somma10anni := età1 + età2
	media10anni := float64(età1+età2) / 2

	fmt.Println("Somma età a 10 anni =", somma10anni)
	fmt.Println("Media età a 10 anni =", media10anni)
}
