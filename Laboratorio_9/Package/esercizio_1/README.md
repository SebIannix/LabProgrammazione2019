# Qual è l'output?

Si considerino le directory:
```text
./
./rettangolo/
```
in cui sono sono presenti i file:
```text
./area.go
./rettangolo/rettangolo.go
```

Nel file `rettangolo.go` è definito il seguente package `rettangolo`:   
```go
package rettangolo

type Rettangolo struct {
	Base, Altezza float64
}

/* Restituisce una nuova istanza del tipo 'Rettangolo' inizializzata
in base ai valori dei parametri 'base' e 'altezza'. */
func NuovoRettangolo(base, altezza float64) *Rettangolo {
	return &Rettangolo{base, altezza}
}

/* Riceve in input un'instanza del tipo `Rettangolo` nel parametro `r` e
restituisce un valore `float64` pari all'area del rettangolo
rappresentato da `r`. */
func Area(r *Rettangolo) float64 {
	return r.Base * r.Altezza
}
```
Supponendo che l'utente inserisca da **riga di comando** i valori 10 e 12, cosa dovrebbe produrre in output il seguente programma scritto nel file `area.go`?

```go
package main

import (
	"./rettangolo"
	"fmt"
	"os"
	"strconv"
)

func main() {
	base, _ := strconv.ParseFloat(os.Args[1], 64)
	altezza, _ := strconv.ParseFloat(os.Args[2], 64)

	r := rettangolo.NuovoRettangolo(base, altezza)

	fmt.Println(rettangolo.Area(r))
}
```
