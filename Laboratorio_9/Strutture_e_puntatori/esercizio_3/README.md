# Trova l'errore

Questo programma dovrebbe stampare `20 100` ma non genera l'output desiderato. Corregere e verificare che l'esecuzione del programma generi l'output atteso.

```go
package main

import "fmt"

func main() {
	var a, b int = 10, 20
	var ptr1, ptr2 *int
	ptr1 = a
	*ptr1 += 10
	*ptr2 = 100
	ptr2 = &b
	fmt.Println(a, b)
}
```
