# Numeri amichevoli

**Definizione**: Due numeri naturali `x` e `y`, con `x < y`, sono detti amichevoli se la somma dei divisori propri di ciascuno è uguale
all’altro; ad esempio `(220, 284)` è una coppia di amichevoli, essendo `284 = 1 + 2 + 4 + 5 + 10 + 11 + 20 + 22 + 44 + 55 + 110` (dove `1, 2, 4, 5, 10, 11, 20, 22, 44, 55, 110` sono i divisori di `220`) e `220 = 1 + 2 + 4 + 71 + 142` (dove `1, 2, 4, 71, 142` sono i divisori di `284`).

Scrivere un programma che legga da **standard input** un numero intero `soglia` e stampi tutte le coppie di numeri amichevoli tali che `y` sia inferiore a `soglia`.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `SommaDivisoriPropri(n int) int` che riceve in input un valore `int` nel parametro `n` e restituisce un valore `int` pari alla somma dei divisori propri di `n` (`0` se `n` non ha divisori propri);
* una funzione `SonoAmichevoli(n, m int) bool` che riceve in input due valori `int` nei parametri `n` ed `m` e restituisce `true` se `n` ed `m` sono amichevoli e `false` altrimenti (utilizzando la funzione `SommaDivisoriPropri()`);
* una funzione `NumeriAmichevoli(limite int)` che riceve in input un valore `int` nel parametro `limite` e stampa tutte le coppie di numeri amichevoli tali che `y` sia inferiore a `limite` (cfr. Definizione); la funzione deve utilizzare la funzione `SonoAmichevoli()`;
* una funzione `main()` che richiama la funzione `LeggiNumero()` per leggere il numero intero `soglia`. Se `soglia > 0` allora richiama la funzione `NumeriAmichevoli()` in modo opportuno, altrimenti stampa un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run numeri_amichevoli.go
Inserisci un numero: 300
Numeri amichevoli a 300
(220,284) 

$ go run numeri_amichevoli.go
Inserisci un numero: 0
La soglia inserita non è positiva
```
