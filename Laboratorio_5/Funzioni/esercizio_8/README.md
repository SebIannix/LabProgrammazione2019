# Numeri primi

Un numero naturale è primo se è divisibile solo per se stesso e per 1.

Scrivere un programma che legga da **standard input** un numero intero `soglia` e stampi tutti i numeri primi inferiori a `soglia`.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `ÈPrimo(n int) bool` che riceve in input un valore (di tipo) `int` nel parametro `n` e restituisce `true` se `n` è primo (i.e., se il valore di `n` rappresenta un numero primo) e `false` altrimenti;
* una funzione `NumeriPrimi(limite int)` che riceve in input un valore (di tipo) `int` nel parametro `limite` e stampa tutti i numeri primi inferiori a `limite` (i.e., al valore di `limite`); la funzione deve utilizzare la funzione `ÈPrimo()`;
* una funzione `main()` che richiama la funzione `LeggiNumero()` per leggere il numero intero `soglia`. Se `soglia > 0` allora richiama la funzione `NumeriPrimi()` in modo opportuno, altrimenti stampa un messaggio d'errore.

##### Esempio d'esecuzione:

```bash
$ go run numeri_primi.go
Inserisci un numero: -3
La soglia inserita non è positiva

$ go run numeri_primi.go
Inserisci un numero: 5
Numeri primi inferiori a 5
2 3 

$ go run numeri_primi.go
Inserisci un numero: 12
Numeri primi inferiori a 12
2 3 5 7 11
```
