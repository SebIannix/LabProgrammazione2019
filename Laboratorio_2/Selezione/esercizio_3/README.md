# Qual è l'output?

Confrontate i programmi che seguono. Cosa producono in output? Il loro funzionamento è identico?

```go

package main

import "fmt"

func main() {

	a := 10
	b := 10

	if a <= b {

		fmt.Println("a <= b")

	} else {

		fmt.Println("a > b")

	}

}
```

```go
package main

import "fmt"

func main() {

	a := 10
	b := 10

	if a <= b {

		fmt.Println("a <= b")

	}
	
	if a > b {

		fmt.Println("a > b")

	}

}
```
