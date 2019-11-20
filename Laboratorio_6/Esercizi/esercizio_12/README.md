# Terna pitagorica

**Definizione**: Se `a`, `b` e `c` sono numeri naturali e `a² + b² = c²`, si dice che la terna di numeri `a`, `b` e `c` è una terna pitagorica (il nome deriva dal teorema di Pitagora: In un triangolo rettangolo di cateti `a`, `b` e di ipotenusa `c` si ha che `a² + b² = c²`). 

Scrivere un programma che legga da **standard input** un intero `n>0` e stampi a video tutte le terne pitagorighe tali che `a<n`, `b<n` e `c<n`. 
 
 Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
 * una funzione `ÈTernaPitagoriga(a int, b int, c int) bool` che riceve in input tre valori `int` nei parametri `a`, `b` e `c`, rispettivamente, e restituisce `true` se `c²` è uguale a `a² + b² ` e `false` altrimenti.

##### Esempio d'esecuzione:
 
```text
$ go run terna_pitagorica.go
Inserisci la soglia: 10
Terne pitagoriche:
(3, 4, 5)
(4, 3, 5)

$ go run terna_pitagorica.go
Inserisci la soglia: 20
Terne pitagoriche:
(3, 4, 5)
(4, 3, 5)
(5, 12, 13)
(6, 8, 10)
(8, 6, 10)
(8, 15, 17)
(9, 12, 15)
(12, 5, 13)
(12, 9, 15)
(15, 8, 17)
```