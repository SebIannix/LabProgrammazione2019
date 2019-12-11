# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import (
	"fmt"
	"strings"
)

type Persona struct {
	Nome, Cognome string
}

func main() {

	p1 := Persona{"Rick", "Sanchez"}

	fmt.Println("Stampe per p1:")
	fmt.Println(p1)
	f(p1)
	fmt.Println(p1)
	g(&p1)
	fmt.Println(p1)

	p2 := NuovaPersona("Mario", "Rossi")

	fmt.Println("\nStampe per p2:")
	fmt.Println(p2)
	f(*p2)
	fmt.Println(p2)
	g(p2)
	fmt.Println(p2)
}

func f(p Persona) {
	p.Nome, p.Cognome = strings.ToUpper(p.Nome), strings.ToUpper(p.Cognome)
}

func g(p *Persona) {
	p.Nome, p.Cognome = strings.ToUpper(p.Nome), strings.ToUpper(p.Cognome)
	// ... equivalente a:
	//(*p).Nome, (*p).Cognome = strings.ToUpper((*p).Nome), strings.ToUpper((*p).Cognome)

}

// Factory method of type Persona
func NuovaPersona(nome, cognome string) (p *Persona) {

	if p == nil {
		fmt.Println("'nil' è lo zero-value di una variabile di tipo " +
			"'reference type' non inizializzata.")
	}
	p = new(Persona) // ... equivalente a: p = &Persona{}
	p.Nome = nome
	p.Cognome = cognome
	return p
	// ... equivalente a:
	//p = &Persona{Nome:nome,Cognome:cognome}
	//return p
	// ... equivalente a:
	//return &Persona{nome, cognome}
}
```
