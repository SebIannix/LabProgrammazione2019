package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	date := LeggiTesto()

	for _, data := range date {
		fmt.Println(Formatta(data))
	}
}

func LeggiTesto() (date []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			return
		}
		date = append(date, riga)
	}
	return
}

func Formatta(data string) string {
	dataRune := []rune(data)

	if len(dataRune) == 8 {
		dataRune = append(dataRune[:4], '-', '0', dataRune[5], '-', '0', dataRune[7])
	} else if len(dataRune) == 9 {
		if dataRune[6] == '/'{
			dataRune = append(dataRune[:4], '-', '0', dataRune[5], '-', dataRune[7], dataRune[8])
		} else {
			dataRune = append(dataRune[:4], '-', dataRune[5], dataRune[6], '-', '0', dataRune[8])
		}
	} else {
		dataRune[4] = '-'
		dataRune[7] = '-'
	}

	return string(dataRune)
}