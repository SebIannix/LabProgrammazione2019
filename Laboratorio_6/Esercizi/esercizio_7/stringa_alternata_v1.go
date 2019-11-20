package main

import (
	"fmt"
)

func main() {

	var s1, s2 string
	fmt.Scan(&s1, &s2)

	fmt.Println(StringheAlternate(s1, s2))

}

func StringheAlternate(s1, s2 string) (risultato string) {

	lunghezzaMassima := len(s1)
	if lunghezzaMassima<len(s2) {
		lunghezzaMassima = len(s2)
	}

	for i := 0; i < lunghezzaMassima; i++ {
    if len(s1) > i {
      risultato += string(s1[i])
    } else {
      risultato += "-"
    }
    if len(s2) > i {
      risultato += string(s2[i])
    } else {
      risultato += "-"
    }
	}

	return
}
