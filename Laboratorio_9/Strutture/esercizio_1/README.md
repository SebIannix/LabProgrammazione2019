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

	var p1, p2, p5 Persona
    // ... equivalente a:
	//p1 = Persona{}
	//p2 = Persona{}
	//p5 = Persona{}

	p1 = Persona{"Rick", "Sanchez"}

	p2.Nome = "Mario"
	p2.Cognome = "Rossi"

	p3 := Persona{"Morty", "Smith"}
	p4 := Persona{Cognome: "Bianchi", Nome: "Paolo"}

	p6 := Persona{Cognome: "Rossi"}

	fmt.Println("p1:", p1)
	fmt.Println("p5:", p5)

	nome := "Giuseppe"
	cognome := "Verdi"

	p5 = Persona{Cognome: cognome, Nome: nome}
    // ... equivalente a:
    // p5 = Persona{nome, cognome}

	fmt.Printf("Nome: %-10s - Cognome: %-10s\n", p1.Nome, p1.Cognome)
	fmt.Printf("%T\t%v\n", p3, p3)
	fmt.Printf("%#v\n", p4)
	fmt.Printf("%#v\n", p6)

}
```
