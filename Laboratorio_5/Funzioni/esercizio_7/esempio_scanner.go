package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	stringa := ""
	fmt.Println("Inserisci una o più righe di testo:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		stringaLocale := scanner.Text()
		fmt.Print("stringaLocale == ",stringaLocale,"\n")
		stringa += stringaLocale + "\n"

	}
	fmt.Println("Righe di testo lette: ")
	fmt.Print(stringa)
}