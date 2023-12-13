package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	for {
		var (
			pilihan int
			tanya   string
		)

		fmt.Println("Program Pilihan")
		fmt.Println("1. Input Data")
		fmt.Println("2. Segitiga")
		fmt.Println("3. Kalkulator")
		fmt.Println("4. Counter HP")
		fmt.Println("5. Luas dan Keliling Bangun Datar")
		fmt.Println("6. Keluar")
		fmt.Print("Masukkan Pilihan [1-6]: ")
		fmt.Scan(&pilihan)

		// Clearing buffer after fmt.Scan call
		fmt.Scanln()

		switch pilihan {
		case 1:
			inputData()
		case 2:
			segitiga()
		case 3:
			kalkulator()
		case 4:
			counterHP()
		case 5:
			bangunDatar()
		case 6:
			fmt.Println("Program selesai. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak ada!")
		}

		fmt.Print("Apakah Anda ingin melanjutkan? (y/n): ")
		fmt.Scan(&tanya)

		for tanya != "y" && tanya != "n" {
			fmt.Println("Input tidak valid. Harap masukkan 'y' atau 'n'.")
			fmt.Print("Apakah Anda ingin melanjutkan? (y/n): ")
			fmt.Scan(&tanya)
		}

		if tanya == "n" {
			break
		}
	}
}

func inputData() {
	var (
		nama, npm, kelas, jurusan string
		uas, uts, total            float64
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nama\t: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("NPM\t: ")
	npm, _ = reader.ReadString('\n')
	npm = strings.TrimSpace(npm)

	fmt.Print("Kelas\t: ")
	kelas, _ = reader.ReadString('\n')
	kelas = strings.TrimSpace(kelas)

	fmt.Print("Jurusan\t: ")
	jurusan, _ = reader.ReadString('\n')
	jurusan = strings.TrimSpace(jurusan)

	fmt.Print("UTS\t: ")
	fmt.Scan(&uts)

	fmt.Print("UAS\t: ")
	fmt.Scan(&uas)

	total = (uas + uts) / 2

	fmt.Println("\n------Hasil------")
	fmt.Printf("Nama\t: %s\n", nama)
	fmt.Printf("NPM\t: %s\n", npm)
	fmt.Printf("Kelas\t: %s\n", kelas)
	fmt.Printf("Jurusan\t: %s\n", jurusan)
	fmt.Printf("Nilai Rata-rata\t: %.2f\n", total)
}

func segitiga() {
	var tinggi int

	fmt.Print("Masukkan Tinggi : ")
	fmt.Scan(&tinggi)

	fmt.Println("------Segitiga------")
	for i := 0; i < tinggi; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j+1, " ")
		}
		fmt.Println()
	}
}

func kalkulator() {
	for {
		var (
			operatorInput, tanya string
			angkaPertama, angkaKedua, hasil float64
		)

		fmt.Print("Masukkan operator (+, -, *, /, % untuk melakukan sisa bagi): ")
		fmt.Scan(&operatorInput)

		fmt.Print("Masukkan angka pertama: ")
		fmt.Scan(&angkaPertama)

		fmt.Print("Masukkan angka kedua: ")
		fmt.Scan(&angkaKedua)

		switch operatorInput {
		case "+":
			hasil = angkaPertama + angkaKedua
		case "-":
			hasil = angkaPertama - angkaKedua
		case "*":
			hasil = angkaPertama * angkaKedua
		case "/":
			if angkaKedua != 0 {
				hasil = angkaPertama / angkaKedua
			} else {
				fmt.Println("Pembagian oleh nol tidak diperbolehkan.")
				continue
			}
		case "%":
			if angkaKedua != 0 {
				hasil = angkaPertama - (float64(int(angkaPertama/angkaKedua)) * angkaKedua)
				fmt.Printf("Sisa bagi dari %.0f %% %.0f = %.0f\n", angkaPertama, angkaKedua, hasil)
			} else {
				fmt.Println("Pembagian oleh nol tidak diperbolehkan.")
				continue
			}
		default:
			fmt.Println("Operator tidak valid.")
			continue
		}

		fmt.Printf("Hasil: %.0f\n", hasil)

		fmt.Print("Apakah Anda ingin melakukan operasi lagi? (y/n): ")
		fmt.Scan(&tanya)

		for tanya != "y" && tanya != "n" {
			fmt.Println("Input tidak valid. Harap masukkan 'y' atau 'n'.")
			fmt.Print("Apakah Anda ingin melakukan operasi lagi? (y/n): ")
			fmt.Scan(&tanya)
		}

		if tanya != "y" && tanya != "Y" {
			break
		}
	}
}

func counterHP() {
	for {
		var (
			pilihanHP    int
			pilihan      string
			harga, bayar float64
		)

		fmt.Println("Counter HP")
		fmt.Println("1. Samsung")
		fmt.Println("2. Infinix")
		fmt.Println("3. Vivo")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Masukkan Pilihan [1-4]: ")
		fmt.Scan(&pilihanHP)

		// Clearing buffer after fmt.Scan call
		fmt.Scanln()

		switch pilihanHP {
		case 1:
			pilihan = "Samsung"
			harga = 3000000
		case 2:
			pilihan = "Infinix"
			harga = 2000000
		case 3:
			pilihan = "Vivo"
			harga = 2500000
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
			continue
		}

		fmt.Printf("Harga %s: %.0f\n", pilihan, harga)

		fmt.Print("Masukkan jumlah uang: ")
		fmt.Scan(&bayar)

		if bayar < harga {
			fmt.Println("Uang yang dimasukkan kurang!")
			continue
		}

		kembalian := bayar - harga
		fmt.Printf("Kembalian: %.0f\n", kembalian)

		fmt.Println("Terima kasih telah berbelanja di Counter HP!")

		fmt.Print("Apakah Anda ingin memilih HP lagi? (y/n): ")
		fmt.Scan(&pilihan)

		for pilihan != "y" && pilihan != "n" {
			fmt.Println("Input tidak valid. Harap masukkan 'y' atau 'n'.")
			fmt.Print("Apakah Anda ingin memilih HP lagi?(y/n): ")
			fmt.Scan(&pilihan)
		}

		if pilihan != "y" && pilihan != "Y" {
			break
		}
	}
}

func bangunDatar() {
	for {
		var (
			pilihanBangun int
			sisiA, sisiB, radius float64
			hasilLuas, hasilKeliling float64
		)

		fmt.Println("Luas dan Keliling Bangun Datar")
		fmt.Println("1. Segitiga")
		fmt.Println("2. Lingkaran")
		fmt.Println("3. Persegi")
		fmt.Println("4. Persegi Panjang")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Masukkan Pilihan [1-5]: ")
		fmt.Scan(&pilihanBangun)

		// Clearing buffer after fmt.Scan call
		fmt.Scanln()

		switch pilihanBangun {
		case 1:
			fmt.Print("Masukkan Panjang Sisi A: ")
			fmt.Scan(&sisiA)
			fmt.Print("Masukkan Panjang Sisi B: ")
			fmt.Scan(&sisiB)

			hasilLuas = luasSegitiga(sisiA, sisiB)
			hasilKeliling = kelilingSegitiga(sisiA, sisiB)
		case 2:
			fmt.Print("Masukkan Radius Lingkaran: ")
			fmt.Scan(&radius)

			hasilLuas = luasLingkaran(radius)
			hasilKeliling = kelilingLingkaran(radius)
		case 3:
			fmt.Print("Masukkan Panjang Sisi Persegi: ")
			fmt.Scan(&sisiA)

			hasilLuas = luasPersegi(sisiA)
			hasilKeliling = kelilingPersegi(sisiA)
		case 4:
			fmt.Print("Masukkan Panjang Persegi Panjang: ")
			fmt.Scan(&sisiA)
			fmt.Print("Masukkan Lebar Persegi Panjang: ")
			fmt.Scan(&sisiB)

			hasilLuas = luasPersegiPanjang(sisiA, sisiB)
			hasilKeliling = kelilingPersegiPanjang(sisiA, sisiB)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
			continue
		}

		fmt.Printf("Luas: %.0f\n", hasilLuas)
		fmt.Printf("Keliling: %.0f\n", hasilKeliling)
	}
}

// Functions for calculating area and perimeter of different shapes
func luasSegitiga(a, b float64) float64 {
	return 0.5 * a * b
}

func kelilingSegitiga(a, b float64) float64 {
	c := math.Sqrt(a*a + b*b)
	return a + b + c
}

func luasLingkaran(r float64) float64 {
	return math.Pi * r * r
}

func kelilingLingkaran(r float64) float64 {
	return 2 * math.Pi * r
}

func luasPersegi(s float64) float64 {
	return s * s
}

func kelilingPersegi(s float64) float64 {
	return 4 * s
}

func luasPersegiPanjang(a, b float64) float64 {
	return a * b
}

func kelilingPersegiPanjang(a, b float64) float64 {
	return 2 * (a + b)
}
