# Conta cifre e lettere

Scrivere un programma che: 
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* stampi a video le seguenti informazioni relative al testo letto:
  1. La lista di lettere distinte presenti nel testo, riportando per ogni lettera il relativo numero di occorrenze nel testo (cfr. **Esecuzione d'esecuzione**). Una lettera è un carattere il cui codice Unicode, se passato come argomento alla funzione `func IsLetter(r rune) bool` del package `unicode`, fa restituire `true` alla funzione. Le lettere minuscole sono da considerarsi diverse dalle lettere maiuscole.
  2. Il numero totale di lettere presenti nel testo.
  3. La lista delle cifre decimali distinte presenti nel testo, riportando per ogni cifra decimale il relativo numero di occorrenze nel testo (cfr. **Esecuzione d'esecuzione**). Una cifra decimale è un carattere il cui codice Unicode, se passato come argomento alla funzione `func IsDigit(r rune) bool` del package `unicode`, fa restituire `true` alla funzione.
  4. Il numero totale di cifre decimali presenti nel testo.

##### Esempio d'esecuzione:

```text
$ go run esercizio_3.go 
Via Celoria 18, 
20133 - Milano
 
REPORT:
Totale lettere: 16
V -> 1
a -> 3
C -> 1
e -> 1
l -> 2
o -> 2
r -> 1
n -> 1
i -> 3
M -> 1
Totale cifre: 7
1 -> 2
8 -> 1
2 -> 1
0 -> 1
3 -> 2

$ go run esercizio_3.go 
Via Celoria 18
REPORT:
Totale lettere: 10
a -> 2
C -> 1
e -> 1
l -> 1
o -> 1
r -> 1
V -> 1
i -> 2
Totale cifre: 2
1 -> 1
8 -> 1

$ go run esercizio_3.go 
44 gatti 
in fila per 3

col resto di 2
REPORT:
Totale lettere: 24
i -> 4
l -> 2
p -> 1
e -> 2
r -> 2
s -> 1
d -> 1
t -> 3
n -> 1
f -> 1
c -> 1
o -> 2
g -> 1
a -> 2
Totale cifre: 4
4 -> 2
3 -> 1
2 -> 1
```
