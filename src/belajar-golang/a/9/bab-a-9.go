// A.9. Tipe Data
package main

import "fmt"

func main() {
	// A.9.1. Tipe Data Numerik Non-Desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// A.9.2. Tipe Data Numerik Desimal
	var decimalNumber = 2.679

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// A.9.3. Tipe Data bool (Boolean)
	var exist bool = false
	fmt.Printf("exist? %t \n", exist)

	// A.9.4. Tipe Data string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)
}
