// A.11. Operator

package main

import (
	"fmt"
)

func main() {
	// A.11.1. Operator Aritmatika
	var value = (((2+6)%3)*4 - 2) / 3
	fmt.Println("hasilnya : ", value)

	// A.11.2. Operator Perbandingan
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	// A.11.3. Operator Logika
	var left = false
	var right = false

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
