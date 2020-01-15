package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	base, _ := strconv.Atoi(os.Args[1])
	esponente, _ := strconv.Atoi(os.Args[2])

	outputFile, err := os.Create("potenze.txt")
	if err != nil {
		fmt.Println("Errore durante la creazione del file:", err)
		return
	}
	defer outputFile.Close()

	if err = Potenze(outputFile, base, esponente); err != nil {
		fmt.Println("Errore durante la scrittura del file:", err)
	}
}

func Potenze(file *os.File, base, esponente int) (err error) {
	potenza := 1
	for i := 0; i <= esponente; i++ {
		if _, err = fmt.Fprintf(file, "%d ", potenza); err != nil {
			return err
		}
		potenza *= base
	}
	return
}
