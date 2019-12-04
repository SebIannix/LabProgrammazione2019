package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	occorrenze := map[rune]int{}

	scanner := bufio.NewScanner(os.Stdin)
	testo := ""
	for scanner.Scan() {
		testo += scanner.Text()
	}

	for _, l := range testo {
		if unicode.IsLetter(l) {
			occorrenze[l]++
		}
	}

	fmt.Println("Occorrenze:")
	for k, v := range occorrenze {
		fmt.Printf("%c: %s\n", k, strings.Repeat("*", v))
	}
}
