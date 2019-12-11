# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

type Persona struct {
	Nome    string
	Cognome string
}

func main() {

	anagrafica := []Persona{
		{"Rick", "Sanchez"},
		{"Morty", "Smith"},
		{"Jerry", "Smith"},
		{"Summer", "Smith"},
		{"Beth", "Smith"}}

	fmt.Printf("Tipo anagrafica %T\n", anagrafica)
	fmt.Println()

	for _, p := range anagrafica {
		fmt.Println(p.Nome, p.Cognome)
	}
	fmt.Println()

	CambiaPrimoCognome(anagrafica)

	for _, p := range anagrafica {
		fmt.Println(p.Nome, p.Cognome)
	}
}

func CambiaPrimoCognome(a []Persona){
	a[0].Cognome = "Martinez"
}
```
