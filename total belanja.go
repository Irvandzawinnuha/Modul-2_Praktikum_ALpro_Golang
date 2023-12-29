package main

import "fmt"

func main() {
	var tahunLahir, totalBelanja int
	fmt.Print("Masukkan tahun lahir: ")
	fmt.Scanln(&tahunLahir)
	fmt.Print("Masukkan total belanja: ")
	fmt.Scanln(&totalBelanja)

	diskon := (tahunLahir % 100) * (tahunLahir / 100) * totalBelanja / 100
	jumlahBayar := totalBelanja - diskon

	fmt.Printf("Besar diskon: %d%%\n", (tahunLahir%100)*(tahunLahir/100))
	fmt.Printf("Jumlah yang dibayar: %d\n", jumlahBayar)
}
