package main

import (
	"fmt"
)

func main() {

	insiemeResiduo := ""

	fmt.Scan(&insiemeResiduo)

	permutazioni := Permutazioni(insiemeResiduo)
	fmt.Printf("%d permutazioni: %v\n", len(permutazioni), permutazioni)

}

func Permutazioni(insiemeResiduo string) (permutazioni []string) {
	if len(insiemeResiduo) == 0 {
		return []string{""}
	} else {
        // L'idea di questa funzione ricorsiva è di vedere la permutazione in questo modo:
        // la permutazione di ABC è formata da tutte le stringhe in cui:
        // - ad A concateniamo ogni permutazione di BC
        // - a B concateniamo ogni permutazione di AC
        // - a C concateniamo ogni permutazione di AB
		for i, v := range insiemeResiduo {
			sottopermutazioni := Permutazioni(insiemeResiduo[:i] + insiemeResiduo[i+1:])
			for _, permutazione := range sottopermutazioni {
				permutazioni = append(permutazioni, string(v) + permutazione)
			}
		}
		return
	}
}
