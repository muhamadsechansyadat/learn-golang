// A.12. Seleksi Kondisi

package main

import (
	"fmt"
)

func main() {
	// A.12.1. Seleksi Kondisi Menggunakan Keyword if, else if, & else
	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	// A.12.2. Variabel Temporary Pada if - else
	var point1 = 8840.0

	if percent := point1 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// A.12.3. Seleksi Kondisi Menggunakan Keyword switch - case
	var point2 = 6

	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// A.12.4. Pemanfaatan case Untuk Banyak Kondisi
	var point3 = 6

	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// A.12.5. Kurung Kurawal Pada Keyword case & default
	var point4 = 6

	switch point4 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// A.12.6. Switch Dengan Gaya if - else
	var point5 = 6

	switch {
	case point5 == 8:
		fmt.Println("perfect")
	case (point5 < 8) && (point5 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// A.12.7. Penggunaan Keyword fallthrough Dalam switch
	var point6 = 6

	switch {
	case point6 == 8:
		fmt.Println("perfect")
	case (point6 < 8) && (point6 > 3):
		fmt.Println("awesome")
		fallthrough
	case point6 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// A.12.8. Seleksi Kondisi Bersarang
	var point7 = 10

	if point7 > 7 {
		switch point7 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point7 == 5 {
			fmt.Println("not bad")
		} else if point7 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point7 == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
