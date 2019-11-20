# Scala (1)

Scrivere un programma che legga da **standard input** un numero intero `n`  e, come mostrato nell'**Esempio d'esecuzione**, stampi a video una scala utilizzando il carattere `'*'` (asterisco):

* La scala è formata da `n` gradini.
* Ciascun gradino è profondo 3 caratteri e alto 2.
* Il gradino più in alto deve essere stampato senza alcuna rientranza verso destra; considerando poi i successivi gradini dall'alto verso il basso, ciascun di essi rientra (è traslato) verso destra rispetto al precedente di due caratteri `' '` (spazio) .

Se `n` è negativo o nullo, anziché stampare la scala il programma deve stampare un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run scala_1.go
4
***
  *
  ***
    *
    ***
      *
      ***
        *

$ go run scala.go
1
***
  *

$ go run scala.go
0
Dimensione non sufficiente
```