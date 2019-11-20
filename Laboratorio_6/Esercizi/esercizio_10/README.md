# Numeri abbondanti

**Definizione**: Un numero naturale è abbondante se è inferiore alla somma dei suoi divisori propri (per esempio, `12` è abbondante perché  la somma dei suoi divisori propri (`1`, `2`, `3`, `4`, `6`) è `16`).

Scrivere un programma che legga da **standard input** un numero intero `soglia` e stampi tutti i numeri abbondanti inferiori a `soglia`.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `SommaDivisoriPropri(n int) int` che riceve in input un valore `int` nel parametro `n` e restituisce un valore `int` pari alla somma dei divisori propri di `n` (pari a `0` se `n` non ha divisori propri);
* una funzione `ÈAbbondante(n int) bool` che riceve in input un valore `int` nel parametro `n` e restituisce `true` se `n` è abbondante, `false` altrimenti; la funzione deve utilizzare la funzione `SommaDivisoriPropri()`;
* una funzione `NumeriAbbondanti(limite int)` che riceve in input un valore `int` nel parametro `limite` e stampa tutti i numeri abbondanti inferiori a `limite`; la funzione deve utilizzare la funzione `ÈAbbondante()`;
* una funzione `main()` che richiama la funzione `LeggiNumero()` per leggere il numero intero `soglia`. Se `soglia > 0` allora richiama la funzione `NumeriAbbondanti()` in modo opportuno, altrimenti stampa un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run numeri_abbondanti.go
20
Numeri abbondanti: 12 18

$ go run numeri_abbondanti.go 
10
Numeri abbondanti:


$ go run numeri_abbondanti.go
30
Numeri abbondanti: 12 18 20 24
```