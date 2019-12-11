# Riduzione di una frazione ai minimi termini

Sulla base del codice sviluppato relativamente all'Esercizio 5 (Laboratorio 9 - Strutture e puntatori), scrivere un programma che:
* legga da **riga di comando**  due numeri interi `n` e `d`;
* riduca ai minimi termini la frazione che ha per numeratore e denominatore i valori `n` e `d`; 
* stampi a video, nel formato `NUMERATORE/DENOMINATORE`, la frazione ridotta ai minimi termini.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Riduci(f *Frazione)` che riceve in input un'instanza del tipo `Frazione` nel parametro `f` e, se necessario, modifica opportunamente il valore dei campi `f.numeratore` e `f.denominatore` affinché `f` rappresenti una frazione ridotta ai minimi termini.

##### Esempio d'esecuzione:
```text
$ go run riduci.go 10 10
1/1

$ go run riduci.go 34 18
17/9

$ go run riduci.go 12 36
1/3
```