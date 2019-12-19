# Scala di numeri

Scrivere un programma che legga da **standard input** due numeri interi `n` e `p` e, come mostrato nell'**Esempio d'esecuzione**, stampi a video una scala utilizzando in sequenza i caratteri corrispondenti alle cifre decimali:

* La scala è formata da `n` gradini.
* Ciascun gradino è profondo `p` caratteri e alto 2.
* Il gradino più in alto deve essere stampato senza alcuna rientranza verso destra; considerando poi i successivi gradini dall'alto verso il basso, ciascuno di essi rientra (è traslato) verso destra rispetto al precedente di `p-1` caratteri `' '` (spazio).
* La sequenza di caratteri utilizzata per rappresentare la scala parte dal carattere `0`. Il carattere successivo al carattere `0` è il carattere `1`. Il carattere successivo al carattere `1` è il carattere `2`, e così via. Il carattere successivo al carattere `9` è il carattere `0`.

Si assuma che vengano sempre inseriti da **standard input** valori interi positivi, sia per `n` sia per `p`.

##### Esempio d'esecuzione

```text
$ go run esercizio_2.go 
1 4
0123
   4

$ go run esercizio_2.go
4 3
012
  3
  456
    7
    890
      1
      234
        5

$ go run esercizio_2.go
3 4
0123
   4
   5678
      9
      0123
         4
```