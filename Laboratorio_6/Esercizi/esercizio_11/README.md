# Statistiche testo

Scrivere un programma che: 
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* stampi a video le seguenti statistiche relative al testo letto:
  1. Il numero di parole presenti nel testo (una parola è una stringa interamente definita da caratteri il cui codice Unicode, se passato come argomento alla funzione `func IsLetter(r rune) bool`, fa restituire `true` alla funzione).
  2. La lunghezza media delle parole presenti nel testo (misurata in caratteri).

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`)) e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `StatisticheParole(s string) (int, int)` che riceve in input un valore `string` nel parametro `s` e restituisce due valori `int` pari rispettivamente al numero di parole presenti in `s` e alla loro lunghezza totale;
* una funzione `main()` che richiama la funzione `LeggiTesto()` per leggere il testo da **standard input**, calcola le statistiche relative al testo letto (richiamando opportunamente la funzione `StatisticheParole()`) e stampa a video le statistiche calcolate.

##### Esempio d'esecuzione:

```text
$ go run statistiche.go
Inserisci un testo su più righe (termina con Ctrl+D):
Testo di prova 

su cui calcolare le statistiche.
:-)
Statistiche:
Numero parole: 8
Lunghezza media: 0.20512820512820512
```
