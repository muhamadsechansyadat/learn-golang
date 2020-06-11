// A.14. Array

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println("names : ", names[0], names[1], names[2], names[3])

	// A.14.1. Pengisian Elemen Array yang Melebihi Alokasi Awal
	var name [4]string
	name[0] = "trafalgar"
	name[1] = "d"
	name[2] = "water"
	name[3] = "law"
	// names[4] = "ez" // baris kode ini menghasilkan error

	// A.14.2. Inisialisasi Nilai Awal Array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// A.14.3. Inisialisasi Nilai Array Dengan Gaya Vertikal
	var fruith [4]string
	var fruitv [4]string

	// cara horizontal
	fruith = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	fruitv = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}
	fmt.Println("Array Horizontal : \t", fruith)
	fmt.Println("Array Vertical : \t", fruitv)

	// A.14.4. Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	var numbers = [...]int{2, 3, 2, 4, 3, 4}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// A.14.5. Array Multidimensi
	var numbers1 = [2][3]string{[3]string{"Putih", "Merah", "Putih"}, [3]string{"Putih", "Merah", "Kuning"}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// A.14.6. Perulangan Elemen Array Menggunakan Keyword for
	var fruit1 = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruit1); i++ {
		fmt.Printf("elemen fruit : %d->%s\n", i, fruit1[i])
	}

	tst := "string"
	tst2 := 10
	tst3 := 1.2

	fmt.Println(reflect.TypeOf(tst))
	fmt.Println(reflect.TypeOf(tst2))
	fmt.Println(reflect.TypeOf(tst3))
}
