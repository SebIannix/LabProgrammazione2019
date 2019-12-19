package main

import "fmt"
import "unicode"
import "bufio"
import "os"

func main() {

    testo := LeggiTesto()
    totLettere, occorrenzeLettere := ContaLettere(testo)
    totCifre, occorrenzeCifre := ContaCifre(testo)

    fmt.Println("REPORT:")
    fmt.Println("Totale lettere:", totLettere)
    for carattere, contatore := range occorrenzeLettere {
        fmt.Printf("%c -> %d\n", carattere, contatore)
    }

    fmt.Println("Totale cifre:", totCifre)
    for carattere, contatore := range occorrenzeCifre {
        fmt.Printf("%c -> %d\n", carattere, contatore)
    }
}

func LeggiTesto() (testo string) {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        testo += scanner.Text() + "\n"
    }

    return
}

func ContaLettere(testo string) (totaleLettere uint, occorrenze map[rune]uint) {
    occorrenze = make(map[rune]uint)
    for _, carattere := range testo {
        if unicode.IsLetter(carattere) {
            totaleLettere ++
            occorrenze[carattere]++
        }
    }
    return
}

func ContaCifre(testo string) (totaleCifre uint, occorrenze map[rune]uint) {
    occorrenze = make(map[rune]uint)
    for _, carattere := range testo {
        if unicode.IsDigit(carattere) {
            totaleCifre ++
            occorrenze[carattere]++
        }
    }
    return
}
