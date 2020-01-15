package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Errore durante l'apertura del file:", err)
		return
	}
	defer inputFile.Close()

	occorrenzaParole, err := ContaOccorrenze(inputFile)
	if err != nil {
		fmt.Println("Errore durante la lettura del file:", err)
		return
	}

	outputFile, err := os.Create("statistiche.txt")
	if err != nil {
		fmt.Println("Errore durante la creazione del file:", err)
		return
	}
	defer outputFile.Close()

	SalvaStatistiche(outputFile, occorrenzaParole)
}

func ContaOccorrenze(file *os.File) (occorrenze map[string]uint, err error) {
	scanner := bufio.NewScanner(file)
	occorrenze = map[string]uint{}
	for scanner.Scan() {
		riga := scanner.Text()
		parole := strings.Split(riga, " ")
		for _, parola := range parole {
			if len(parola) > 0 {
				occorrenze[strings.ToUpper(parola)]++
			}
		}
	}
	err = scanner.Err()
	return
}

func SalvaStatistiche(file *os.File, occorrenze map[string]uint) (err error) {
	if _, err = fmt.Fprintf(file, "Totale parole distinte: %d\nOccorrenze:\n", len(occorrenze)); err != nil {
		return
	}

	for parola, contatore := range occorrenze {
		if _, err = fmt.Fprintf(file, "%s: %d\n", parola, contatore); err != nil {
			return
		}
	}
	return
}