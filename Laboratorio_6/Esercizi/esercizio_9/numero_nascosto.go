package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	riga := scanner.Text()

	numeroNascosto := ""
	for _, carattere := range riga {
		if unicode.IsDigit(carattere) {
			numeroNascosto += string(carattere)
		}
	}

	numero, _ := strconv.Atoi(numeroNascosto)
	fmt.Print("Doppio del numero nascosto: ", 2*numero, " (", numero, " * 2)\n")

}

