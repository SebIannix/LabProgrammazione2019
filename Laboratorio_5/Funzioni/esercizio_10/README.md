# Numeri perfetti

**Definizione**: Un numero naturale è perfetto se è uguale alla somma dei suoi divisori propri (per esempio, `6` è perfetto perché `6 = 1 + 2 + 3`).

Scrivere un programma che legga da **standard input** un numero intero e stampi se il numero è perfetto oppure no.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `SommaDivisoriPropri(n int) int` che riceve in input un valore (di tipo) `int` nel parametro `n` e restituisce un valore  (di tipo) `int` pari alla somma dei divisori propri di `n` (i.e., alla somma dei divisori propri del numero rappresentato dal valore di `n`), pari a `0` se (il numero rappresentato dal valore di) `n` non ha divisori propri;
* una funzione `ÈPerfetto(n int) bool` che riceve in input un valore  (di tipo) `int` nel parametro `n` e restituisce `true` se `n` è perfetto (i.e., se il valore di `n` rappresenta un numero perfetto) e `false` altrimenti; la funzione deve utilizzare la funzione `SommaDivisoriPropri()`;
* una funzione `main()` che richiama la funzione `LeggiNumero()` per leggere il numero intero `numero` da controllare. Se `numero > 0` allora richiama la funzione `ÈPerfetto()` in modo opportuno, altrimenti stampa un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run numero_perfetto.go
Inserisci un numero: 0
0 non è perfetto

$ go run numero_perfetto.go
Inserisci un numero: 1
1 non è perfetto

$ go run numero_perfetto.go
Inserisci un numero: 6
6 è perfetto

$ go run numero_perfetto.go
Inserisci un numero: 2
2 non è perfetto

$ go run numero_perfetto.go
Inserisci un numero: 28
28 è perfetto
```

 
 