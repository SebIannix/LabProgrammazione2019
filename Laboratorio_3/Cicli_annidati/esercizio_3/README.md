# Qual è l'output?

Quanti asterischi stampa il seguente programma?

```go
package main

import "fmt"

func main() {
	n := 3

	for i := 0; i < n; i++ {
    	for j := i+1; j < n; j++ {
    		fmt.Print("*")
    	}
    }
    fmt.Println()
}
```