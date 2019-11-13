# Garibaldi

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da **standard input** una riga vuota (`""`);
* ristampi il testo letto (riga vuota esclusa) sostituendo tutte le vocali non accentate (minuscole o maiuscole) con delle `u`.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore (di tipo) `string` (una stringa) in cui è memorizzato il testo letto (riga vuota esclusa);
* una funzione `TrasformaCarattere(c rune) rune` che riceve in input un valore (di tipo) `rune` nel parametro `c` e restituisce un valore (di tipo) `rune` uguale a `u` se `c` è una vocale non accentata (minuscola o maiuscola) (i.e., se il valore di `c` rappresenta una vocale non accentata (minuscola o maiuscola)) e uguale a `c` altrimenti.
* una funzione `Garibaldi(s string) string` che riceve in input un valore (di tipo) `string` (una stringa) nel parametro `s` e restituisce un valore (di tipo) `string` (una stringa) in cui l'i-esimo carattere è uguale all'i-esimo carattere di `s` se diverso da una vocale non accentata e uguale a `u` altrimenti; la funzione deve utilizzare la funzione `TrasformaCarattere()`;
* una funzione `main()` che richiama la funzione `LeggiTesto()` per leggere il testo da **standard input**, trasforma il testo letto (richiamando opportunamente la funzione `Garibaldi()`) e stampa il testo trasformato.

##### Esempio d'esecuzione:

```text
$ go run garibaldi.go
Inserisci un testo su più righe (termina con riga vuota):
Garibaldi fu ferito
fu ferito in una gamba,
Garibaldi che comanda
che comanda i bersaglier

Risultato trasformazione:
Gurubuldu fu furutu
fu furutu un unu gumbu,
Gurubuldu chu cumundu
chu cumundu u bursugluur
```
