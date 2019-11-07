package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// inizializza il generatore di numeri casuali
	// richiede l'importazione sia del package math/rand che del package time
	rand.Seed(int64(time.Now().Nanosecond()))

	// rand.Intn(100) genera un numero casuale tra 0 e 99.
	// Aggiungo 1 al numero estratto per generare un numero casuale tra 1 e 100
	numeroEstratto := rand.Intn(100) + 1
	tentativi := 1

	// questo ciclo è senza condizione e può terminare solamente con l'esecuzione dell'istruzione break
	for {
		var numeroLetto int
		fmt.Print("Tentativo n° ", tentativi, ": ")
		fmt.Scan(&numeroLetto)

		// se il numero inserito è uguale al numero estratto allora termino il ciclo for
		if numeroLetto == numeroEstratto {
			break
			// altrimenti comunico se il numero inserito è maggiore o minore di quello estratto
		} else if numeroLetto > numeroEstratto {
			fmt.Println("Troppo alto! Riprova!")
		} else {
			fmt.Println("Troppo basso! Riprova!")
		}

		tentativi++
	}

	fmt.Println("Hai indovinato in ", tentativi, " tentativi!")
}
