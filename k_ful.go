package main

import (
	"fmt"
)

func main() {
	// Input nilai UTS dan UAS
	var uts, uas float64
	fmt.Print("Masukkan nilai UTS: ")
	fmt.Scanln(&uts)
	fmt.Print("Masukkan nilai UAS: ")
	fmt.Scanln(&uas)

	// Menghitung nilai akhir
	nilaiAkhir := (uts + uas) / 2

	// Menentukan keterangan berdasarkan rentang nilai
	var keterangan string
	if nilaiAkhir >= 90 {
		keterangan = "A"
	} else if nilaiAkhir >= 80 {
		keterangan = "B"
	} else if nilaiAkhir >= 70 {
		keterangan = "C"
	} else if nilaiAkhir >= 60 {
		keterangan = "D"
	} else {
		keterangan = "E"
	}

	// Menampilkan hasil
	fmt.Printf("Nilai Akhir: %.2f\n", nilaiAkhir)
	fmt.Printf("Keterangan: Grade %s\n", keterangan)
}
