// A.10. Konstanta

package main

import "fmt"

func main() {
	// A.10.1. Penggunaan Konstanta
	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	const lastName = "wick"
	fmt.Print("nice to meet you ", lastName, "!\n")

	// A.10.1.1. Penggunaan Fungsi fmt.Print()
	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")
}
