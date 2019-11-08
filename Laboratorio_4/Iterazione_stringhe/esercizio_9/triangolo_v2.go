package main

import "fmt"

func main() {
	var dimensione int
	fmt.Scan(&dimensione)

	// Versione 2:
	// uso un ciclo per creare una stringa di *
	// uso un secondo ciclo che stampa ad ogni iterazione una parte della stringa

	if dimensione <= 0 {
		fmt.Println("Dimensione non sufficiente")
	} else {
		// creo una stringa vuota che ad ogni iterazione crescerà
		// aggiungendo un * in fondo alla stringa
		stringaAsterischi := ""
		for i := 0; i < dimensione; i++ {
			stringaAsterischi += "*"
		}

		// Ad ogni iterazione, questo ciclo for stamperà i+1 caratteri della stringa:
		// alla prima iterazione stampa 1 carattere (la fetta stringaAsterischi[0:1])
		// alla seconda iterazione stampa 2 caratteri (la fetta stringaAsterischi[0:2])
		// alla terza iterazione stampa 3 caratteri (la fetta stringaAsterischi[0:3])
		// e così via
		for i := 0; i < dimensione; i++ {
			fmt.Print(stringaAsterischi[:i+1], "\n")
		}
	}

}
