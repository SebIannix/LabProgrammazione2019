# Triangoli casuali

Si estenda il package `triangolo` definito relativamente all'Esercizio 2 (Laboratorio 9 - Package) implementando la funzione:

* `String(t Triangolo) string` che riceve in input un'instanza del tipo `Triangolo` nel parametro `t` e restituisce un valore `string` che corrisponde alla rappresentazione `string` di `t` nel formato `Triangolo con lati L1, L2 e L3.`, dove `L1`, `L2` ed `L3` sono i valori `string` corrispondenti ai valori dei campi `lato1`, `lato2` e `lato3` di `t`.

Utilizzando le funzionalità messe a disposizione dal package `triangolo`, scrivere un programma che:
* legga da **riga di comando** un numero intero `n`;
* generi in maniera casuale `n` triple di valori reali compresi tra `10` e `1000`; i valori `l1`, `l2`, `l3` di ciascuna tripla corrispondono alle misure dei lati di un ipotetico triangolo;
* stampi a video la rappresentazione `string` del triangolo con area più grande tra quelli corrispondenti alle triple di valori reali generate; 
* stampi a video la rappresentazione `string` del triangolo con perimetro più piccolo tra quelli corrispondenti alle triple di valori reali generate.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `GeneraTriangoli(n int) (tN []triangolo.Triangolo)` che riceve in input un valore `int` nel parametro `n` e restituisce un valore `[]triangolo.Triangolo` nella variabile `tN` in cui sono memorizzate le istanze del tipo `triangolo.Triangolo` inizializzate in base alle `n` triple di valori reali generate in maniera casuale (poiché la funzione `triangolo.NuovoTriangolo(l1, l2, l3 float64) (t *triangolo.Triangolo, err bool)` restituisce i valori `nil` e `true` nel caso in cui `l1+l2 <= l3`, `l1+l3 <= l2` o `l2+l3 <= l1`, potrebbe succedere che `len(tN) < n`).


##### Esempio d'esecuzione:

```text
$ go run triangoli_casuali.go 10
Triangolo con area maggiore = Triangolo con lati 844.728464, 872.971432 e 644.031285
Triangolo con perimetro minore = Triangolo con lati 587.135063, 363.401214 e 612.413022
```
