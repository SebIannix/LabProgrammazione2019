# Spaziatura

Scrivere un programma che: 
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* ristampi il testo letto con spaziatura, ovvero inserendo uno spazio `' '` tra ogni coppia di caratteri che **non** sono spazi.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`)) e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `Spazia(s string) string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `string` che rappresenta la versione spaziata di `s`;
* una funzione `main()` che richiama la funzione `LeggiTesto()` per leggere il testo da **standard input**, spazia il testo letto (richiamando opportunamente la funzione `Spazia()`) e stampa il testo spaziato.

*Suggerimento:* per sapere se un carattere è uno spazio si faccia riferimento alla documentazione della funzione `unicode.IsSpace` del package `unicode`.

##### Esempio d'esecuzione:

```text
$ go run spaziatura.go 
Inserisci un testo su più righe (termina con Ctrl+D):
Testo di prova 
da stampare con spaziatura
Risultato:
T e s t o d i p r o v a
d a s t a m p a r e c o n s p a z i a t u r a
```
