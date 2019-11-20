# Robot

Un robot è in grado di muoversi nelle quattro direzioni nord (N), sud (S), est (E) e ovest (O). 

In particolare, il robot accetta comandi che consistono in una coppia di valori `d p`, dove `d` è un carattere che indica la direzione e può assumere uno dei quattro valori `'N'`, `'S'`, `'E'`, `'O'`, mentre `0 < p < 10` è un intero che indica il numero di passi che il robot deve compiere in quella direzione. 

Scrivere un programma che, dopo aver letto da **standard input** un numero intero `n`, chiede all'utente di inserire `n` comandi da impartire al robot (sempre da **standard input**). 

Dopo aver letto gli `n` comandi, il programma deve stampare a video:
* il numero totale di passi che deve compiere il robot in ognuna delle quattro direzioni (se il robot non si deve mai muovere in una certa direzione, l'output relativo a tale direzione non va stampato);
* la sequenza di comandi inversi che si dovrebbe impartire al robot per farlo ritornare al punto di partenza lungo lo stesso percorso.

*Suggerimento*: La sequenza di comandi:
```text
O 2
O 3 
N 4 
```
può essere memorizzata nella stringa `"O2O3N4"`. In particolare, una qualsiasi sequenza di comandi può essere memorizzata in una stringa interamente definita da caratteri considerati nello standard US-ASCII.

##### Esempio d'esecuzione:

```text
$ go run robot.go
Inserisci il numero di comandi:
5
Inserisci i comandi:
O 2 
O 3 
N 4 
S 5 
E 1 

Movimenti totali:
N 4
S 5
E 1
O 5
Comandi inversi:
O 1, N 5, S 4, E 3, E 2

$ go run robot.go 
Inserisci il numero di comandi:
8
Inserisci i comandi:
O 2  
N 1
E 5 
S 4 
E 5 
N 3 
S 1
O 3 

Movimenti totali:
N 4
S 5
E 10
O 5
Comandi inversi:
E 3, N 1, S 3, O 5, N 4, O 5, S 1, E 2

$ go run robot.go  
Inserisci il numero di comandi:
4
Inserisci i comandi:
E 3
N 1
N 4
E 2

Movimenti totali:
N 5
E 5
Comandi inversi:
O 2, S 4, S 1, O 3

$ go run robot.go 
Inserisci il numero di comandi:
5
Inserisci i comandi:
N 3 
N 2 
O 3 
S 2 
O 1 

Movimenti totali:
N 5
S 2
O 4
Comandi inversi:
E 1, N 2, E 3, S 2, S 3
```
