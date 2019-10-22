# Area poligono regolare

Scrivere un programma che legga da **standard input** due numeri interi che chiameremo `n` e `l` e calcoli l’area di un poligono regolare con `n` lati di lunghezza `l`.

*Suggerimento:* l'area di un poligono regolare può essere calcolata utilizzando le funzioni `math.Pow` (per il calcolo della potenza) e `math.Tan` (per il calcolo della tangente di un angolo) del package `math` nel seguente modo:
```go
area := (n * math.Pow(l, 2)) / (4 * math.Tan(math.Pi/n))
```
