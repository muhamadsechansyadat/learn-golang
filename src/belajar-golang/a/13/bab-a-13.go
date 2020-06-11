// A.13. Perulangan

package main

import "fmt"

func main() {
	// A.13.1. Perulangan Menggunakan Keyword for
	for i := 1; i <= 5; i++ {
		fmt.Println("Angka : ", i)
	}

	// A.13.2. Penggunaan Keyword for Dengan Argumen Hanya Kondisi
	var i = 1

	for i <= 5 {
		fmt.Println("Angka-i-1 : ", i)
		i++
	}

	// A.13.3. Penggunaan Keyword for Tanpa Argumen
	var a = 0

	for {
		fmt.Println("Angka-a : ", a)

		a++
		if a == 5 {
			break
		}
	}

	// A.13.5. Penggunaan Keyword break & continue
	for b := 1; b <= 10; b++ {
		if b%2 == 1 {
			continue
		}

		if b > 8 {
			break
		}

		fmt.Println("Angka-b", b)
	}

	// A.13.6. Perulangan Bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// A.13.7. Pemanfaatan Label Dalam Perulangan
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
