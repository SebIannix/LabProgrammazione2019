# Istogramma a barre orizzontali (2)

Scrivere un programma che: 
1. legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
2. termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
3. come mostrato nell'**Esempio di esecuzione**, stampi un istogramma a barre orizzontali per rappresentare il numero di occorrenze di ogni lettera presente nel testo letto:
* una lettera è un carattere il cui codice Unicode, se passato come argomento alla funzione `func IsLetter(r rune) bool` del package `unicode`, fa restituire `true` alla funzione;
* le lettere minuscole sono da considerarsi diverse dalle lettere maiuscole; 
* ogni barra viene rappresentata utilizzando il carattere asterisco (`*`); se il numero di occorrenze della lettera `e` è per esempio `9`, la barra corrispondente sarà formata da `9` caratteri `*`;
* le barre devono essere stampate a partire da quella associata alla lettera con codice Unicode più piccolo fino a quella associata alla lettera con codice Unicode più grande.    

*Suggerimento*: 
* Si consideri il seguente programma.

```go
package main

import (
	"fmt"
)

func main() {

	capitali := map[string]string{"Francia": "Parigi", "Italia": "Roma", 
        "Giappone": "Tokio", "Austria": "Vienna"}

	stati := []string{"Austria", "Francia", "Giappone", "Italia"}

	for _, v := range stati {
		fmt.Println(capitali[v])
	}
}
```
Output:
```text 
Vienna
Parigi
Tokio
Roma
```
Il programma stampa a video i nomi delle capitali europee memorizzati in `capitali` in base all'ordine in cui i nomi degli stati corrispondenti sono memorizzati in `stati`.
* La seguente funzione ordina una slice di rune in ordine crescente.

```go
func SortRunes(a []rune) {
	for i:=0;i<len(a)-1;i++{
		indiceMin := i
		for j := i + 1; j<len(a); j++ {
			if a[indiceMin] > a[j] {
				indiceMin = j
			}
		}
		a[i], a[indiceMin] = a[indiceMin], a[i]
	}
}
```

##### Esempio d'esecuzione:

```text
$ go run istogramma.go
Ciao,
come stai?
Occorrenze:
C: *
a: **
c: *
e: *
i: **
m: *
o: **
s: *
t: *
```