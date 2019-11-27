# Qual è l'outout?

Qual è l'output di questo programma?

```go
package main

import "fmt"
import "math"

func main() {
    
    var n int = 5

    //a := [n]int   /* non corretto */ 
 
    sl := make([]int, n)

    for i:=0;i<len(sl);i++{
         sl[i] = i
    }      

    fmt.Printf("%T: %v\n", sl, sl)
}
```