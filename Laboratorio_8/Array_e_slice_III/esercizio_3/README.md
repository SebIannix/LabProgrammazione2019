# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

const Dimensione = 10

func main() {

	var a []int

	for i := 0; i < Dimensione; i++ {
		a = append([]int{i+1}, a...)
	}

	fmt.Println(a)

	b := make([]int, Dimensione)

	copy(b[:], a[Dimensione/2:])

	fmt.Println(b)

	c := make([]int, Dimensione)

	for i := 0; i < Dimensione; i += 1 {
		copy(c[Dimensione-1-i:], a[i:i+1])
	}
	fmt.Println(c)
}
```