package main

import (
	"fmt"
)

func main() {

	insiemeResiduo := ""

	fmt.Scan(&insiemeResiduo)

	sottoinsiemi := Sottoinsiemi(insiemeResiduo)
	fmt.Printf("%d sottoinsiemi: %v\n", len(sottoinsiemi), sottoinsiemi)

}

func Sottoinsiemi(insiemeResiduo string) (sottoinsiemi []string) {
	if len(insiemeResiduo) == 0 {
		return
	} else {
        // L'idea ricorsiva di questa funzione è quella di formare i sottoinsiemi in questo modo:
        // dato l'insieme residuo ABC, calcolo i sottoinsiemi di BC
        // Per ogni sottoinsieme di BC, creo una copia di tale sottoinsieme in cui aggiungo anche A
        // In questo modo partiziono i sottoinsiemi tra quelli che hanno l'elemento A e quelli che non ce l'hanno
		sottoinsiemi = Sottoinsiemi(insiemeResiduo[1:])
		for _, sottoinsieme := range sottoinsiemi {
			sottoinsiemi = append(sottoinsiemi, string(insiemeResiduo[0]) + sottoinsieme)
		}
		sottoinsiemi = append(sottoinsiemi, string(insiemeResiduo[0]))
		return
	}
}
