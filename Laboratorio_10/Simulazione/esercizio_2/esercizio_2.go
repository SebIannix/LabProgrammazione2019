package main

import "fmt"

func main() {

    var n, p int
    fmt.Scan(&n, &p)

    StampaScala(n, p)

}

func StampaScala(gradini, profondità int) {
    contatore := 0 
    for i:=0; i<gradini; i++ {
        StampaGradino(i, profondità, &contatore)    
    }
}

func StampaGradino(numeroGradino, profondità int, contatore *int) {
    for i:=0; i<numeroGradino*(profondità-1); i++ {
        fmt.Print(" ")
    }

    for i:=0; i<profondità; i++ {
        fmt.Print(*contatore)
        *contatore = (*contatore+1)%10
    }
    fmt.Println()

    for i:=0; i<numeroGradino*(profondità-1); i++ {
        fmt.Print(" ")
    }
    for i:=0; i<profondità-1; i++ {
        fmt.Print(" ")
    }
    fmt.Println(*contatore)
    *contatore = (*contatore+1)%10
}
