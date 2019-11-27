# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	a := [...]string{5:"E", 3:"C"}

	fmt.Println(a)

	for i:=len(a)-1; i>=0; i-- {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}
	fmt.Println()

	for i, _ := range a {
		fmt.Println("Indice", len(a)-1-i, " - Valore:", a[len(a)-1-i])
	}
	fmt.Println()

	for i, _ := range a {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}
	fmt.Println()

	for i := range a {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}

}
```