package main

import "fmt"

func main() {
	var angka int
    // Meminta pengguna memasukkan bilangan bulat
    fmt.Print("Masukkan sebuah bilangan bulat: ")
    fmt.Scan(&angka)

    // Menentukan apakah bilangan positif, negatif, atau nol
    switch {
    case angka > 0:
        fmt.Println("Bilangan tersebut adalah positif.")
    case angka < 0:
        fmt.Println("Bilangan tersebut adalah negatif.")
    default:
        fmt.Println("Bilangan tersebut adalah nol.")
    }
}