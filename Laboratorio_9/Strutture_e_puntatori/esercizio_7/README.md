# Moltiplicazione tra frazioni

Sulla base del codice sviluppato relativamente all'Esercizio 5 e 6 (Laboratorio 9 - Strutture e puntatori), scrivere un programma che:
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da standard input una riga vuota (`""`).

In ogni riga sono specificati due numeri interi che rappresentano il numeratore e il denominatore di una frazione.

Una volta terminata la fase di lettura, il programma deve:
1. calcolare la frazione che si ottiene moltiplicando tra di loro le frazioni corrispondenti alle coppie di interi letti;
2. ridurre ai minimi termini la frazione calcolata al punto 1;
3. stampare a video, nel formato `NUMERATORE/DENOMINATORE`, la frazione ridotta ai minimi termini ottenuta al punto 2.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `LeggiFrazioni() []Frazione` che legge da **standard input**  un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore `[]Frazione` in cui sono memorizzate tutte le istanze del tipo `Frazione` inizializzate in base ai numeri interi `numeratore` e `denominatore` specificati in ciascuna delle righe lette;
* una funzione `Moltiplica(f1, f2 Frazione) *Frazione` che riceve in input due istanze del tipo `Frazione` nei parametri `f1` e `f2` e restituisce una nuova istanza del tipo `Frazione` i cui campi `numeratore` e `denominatore` sono inizializzati con i valori  `f1.numeratore * f2.numeratore` e `f1.denominatore * f2.denominatore`;
* una funzione `MoltiplicaN(fN []Frazione) *Frazione` che riceve in input un valore `[]Frazione` nel parametro `fN` e restituisce una nuova istanza del tipo `Frazione` corrispondente alla frazione che si ottiene moltiplicando tra di loro le frazioni relative alle istanze del tipo `Frazione` presenti in `fN`; la funzione deve utilizzare la funzione `Moltiplica()`.

##### Esempio d'esecuzione:
```text
$ go run prodotto.go 
Inserisci numeratore e denominatore delle frazioni:
1 2
4 3
2 5

4/15

$ go run prodotto.go
Inserisci numeratore e denominatore delle frazioni:
1 1
2 3
1 8

1/12
```
