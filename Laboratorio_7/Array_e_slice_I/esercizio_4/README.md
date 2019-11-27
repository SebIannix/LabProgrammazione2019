# Qual è l'outout?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a [6]int
    
    for i:=0;i<len(a);i++{
        a[i] = i*2
    }   

    fmt.Printf("a - %T: %v\n",a,a)

    b := a

    fmt.Printf("b - %T: %v\n",b,b)

	sl1 := a[:]    
    sl2 := a[2:]

    fmt.Printf("\na - %v\n",a)
    fmt.Printf("sl1 - %T: %v\n",sl1,sl1)

    sl1[0] = -9
    sl2[0] = 100

    fmt.Printf("a - %v\n",a)

    fmt.Printf("\nb - %v\n",b)

	for _, v := range b {
		v *= 2
	}
    
    fmt.Printf("b - %v\n",b)
}
```