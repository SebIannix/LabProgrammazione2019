package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	dividendo, divisore := LeggiValori()
	quoziente, resto := Dividi(dividendo, divisore)
	fmt.Printf("Quoziente = %d, Resto = %d\n", quoziente, resto)

}

func LeggiValori() (a, b int) {
	a, _ = strconv.Atoi(os.Args[1])
	b, _ = strconv.Atoi(os.Args[2])
	return
}

func Dividi(dividendo, divisore int) (quoziente, resto int) {
    if divisore == 0 {
        return 0, 0 // questo caso gestisce l'eventuale errore di una divisione per 0
                    // ma non era richiesto dal testo dell'esercizio
    } else if dividendo < divisore {
		return 0, dividendo
	} else {
		quoziente, resto = Dividi(dividendo - divisore, divisore)
		return quoziente + 1, resto
	}
}
