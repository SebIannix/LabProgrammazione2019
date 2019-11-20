# Numero nascosto

Scrivere un programma che legga da **standard input** una riga di testo e stampi in output il doppio del numero nascosto all'interno della riga di testo, ovvero il doppio del numero che si ottiene concatenando le cifre presenti all'interno della riga di testo.

*Suggerimento:* Per convertire una stringa in un numero intero  si faccia riferimento alla documentazione della funzione  `strconv.Atoi()` del package `strconv`.

##### Esempio d'esecuzione:

```text
$ go run numero_nascosto.go
Ci8ao 97com3 va?
Doppio del numero nascosto: 17946 (8973 * 2)

$ go run numero_nascosto.go
Ch3 831 t3mp0
Doppio del numero nascosto: 766260 (383130 * 2)

$ go run numero_nascosto.go
c140n3
Doppio del numero nascosto: 2806 (1403 * 2)
```