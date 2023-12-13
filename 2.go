package main

import "fmt"

func hitungLuasSegitiga(alas, tinggi float64) float64 {
    return 0.5 * alas * tinggi
}

func main() {
    var (
        alas   float64
        tinggi float64
    )

    fmt.Print("Masukkan alas segitiga: ")
    fmt.Scan(&alas)

    fmt.Print("Masukkan tinggi segitiga: ")
    fmt.Scan(&tinggi)

    luas := hitungLuasSegitiga(alas, tinggi)
    fmt.Printf("Luas Segitiga: %.2f\n", luas)
}
