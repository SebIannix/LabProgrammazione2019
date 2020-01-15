package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputFile *os.File
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Errore aprendo il file", os.Args[1], "-", err)
		return
	}
	defer inputFile.Close()

	pattern := os.Args[2]
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		if riga := scanner.Text(); strings.Contains(riga, pattern) {
			fmt.Println(riga)
		}
	}
}
