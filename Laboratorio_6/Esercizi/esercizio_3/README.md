# Qual è l'output?

Supponendo che l'utente inserisca da **standard input** il valore `5`, cosa dovrebbe produrre in output il seguente programma?

```go
package main

import (
	"fmt"
)

func main() {

	var cifraInInput int32

	for cifraInInput <= 0 || cifraInInput >= 10 {
		fmt.Print("Inserisci un'unica cifra decimale maggiore o uguale a 1: ")
		fmt.Scan(&cifraInInput)
	}
	fmt.Println()

	var carattere rune

	carattere = '0' + (cifraInInput - 1)

	fmt.Printf("Carattere: %c\n", carattere)
    fmt.Printf("Carattere: %d\n", carattere)

	var cifraInOutput int32

	cifraInOutput = (carattere - '0') + 1

	fmt.Printf("Cifra: %d\n", cifraInOutput)
    fmt.Printf("Cifra: %c\n", cifraInOutput)
}
```