package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	testo := LeggiTesto()
	testoInvertito := InvertiStringa(testo)

	fmt.Println("Testo invertito:")
	fmt.Println(testoInvertito)

}

func LeggiTesto() (testo string) {
	fmt.Println("Inserisci un testo su più righe (termina con riga vuota):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga != "" {
			testo += riga + "\n"
		} else {
			break
		}
	}
	return testo
}

func InvertiStringa(stringa string) (stringaOutput string) {
	for _, c := range stringa {
		stringaOutput = string(c) + stringaOutput
	}
	return
}
