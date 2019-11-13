# Numeri primi gemelli

**Definizione**: Due numeri primi `p` e `q` sono gemelli se `p = q + 2`. 

Scrivere un programma che legga da **standard input** un numero intero `soglia` e stampi tutti i numeri primi gemelli tali che `p` sia inferiore a `soglia`.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `ÈPrimo(n int) bool` che riceve in input un valore `int` nel parametro `n` e restituisce `true` se `n` è primo e `false` altrimenti;
* una funzione `NumeriPrimiGemelli(limite int)` che riceve in input un valore `int` nel parametro `limite` e stampa tutte le coppie di numeri primi gemelli tali che `p` sia inferiore a `limite` (cfr. Definizione); la funzione deve utilizzare la funzione `ÈPrimo()`;
* una funzione `main()` che richiama la funzione `LeggiNumero()` per leggere il numero intero `soglia`. Se `soglia > 0` allora richiama la funzione `NumeriPrimiGemelli()` in modo opportuno, altrimenti stampa un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run numeri_primi_gemelli.go
Inserisci un numero: -4
La soglia inserita non è positiva

$ go run numeri_primi_gemelli.go
Inserisci un numero: 10
Numeri primi gemelli inferiori a 10
(3,5)(5,7)

$ go run numeri_primi_gemelli.go
Inserisci un numero: 10
Numeri primi gemelli inferiori a 10
(3,5) (5,7) 

$ go run numeri_primi_gemelli.go
Inserisci un numero: 20
Numeri primi gemelli inferiori a 20
(3,5) (5,7) (11,13) (17,19)
```