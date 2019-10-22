# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {
	var (
		a = 10
		b = 20
		c = 30
	)
	if a > b {
		a = b
	} else {
		b = a
	}
	c = c + b + a
	fmt.Println(a, b, c)
}
```
