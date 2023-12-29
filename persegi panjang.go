package main

import (
	"fmt"
	"math"
)

func main() {
	var panjang, lebar float64
	fmt.Print("Masukkan panjang: ")
	fmt.Scanln(&panjang)
	fmt.Print("Masukkan lebar: ")
	fmt.Scanln(&lebar)

	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	diagonal := math.Sqrt((panjang * panjang) + (lebar * lebar))

	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n", keliling)
	fmt.Printf("Panjang diagonal: %.2f\n", diagonal)
}