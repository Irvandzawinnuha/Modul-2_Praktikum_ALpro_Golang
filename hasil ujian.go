package main

import "fmt"

func main() {
	var n, n_passed, n_failed int
	var n1, n2, n3, avg float64

	fmt.Println("Berapa jumlah siswa yang nilainya akan diproses?")
	fmt.Scan(&n)

	n_passed = 0
	n_failed = 0

	for i := 1; i <= n; i++ {
		fmt.Println("Masukkan nilai siswa", i, "(n1 n2 n3):")
		fmt.Scan(&n1, &n2, &n3)

		avg = (n1 + n2 + n3) / 3.0

		if avg > 80.0 {
			fmt.Println("Memenuhi syarat administratif")
			n_passed++
		} else {
			fmt.Println("Tidak memenuhi syarat administratif")
			n_failed++
		}
	}

	fmt.Println("Jumlah siswa lolos seleksi administrasi:", n_passed)
	fmt.Println("Jumlah siswa tidak lolos seleksi administrasi:", n_failed)
}
