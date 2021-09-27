package main

import "fmt"

func main() {
	// LatihanVariable()
	// LatihanIfControl()
	// LatihanSwitchControl()
	// LatihanForControl()
	// LatihanArray()
	// LatihanSlice()
	// LatihanMap()
	// LatihanSliceMap()
	// Quiz2()

	Quiz3()

	/* BELAJAR FUNCTION */
	// printMyResult("Saya sedang belajar Golang")

	// // luas, keliling := calculate(10, 2)
	// luas, _ := calculate(10, 2)
	// fmt.Println("Luas : ", luas)
	// // fmt.Println(keliling)

	// // luas, keliling := calculate(10, 2)
	// luas, keliling := prCalculate(20, 4)
	// fmt.Println("PR : ", luas)
	// fmt.Println("PR : ", keliling)

}

/* FUNCTION */

/* =================
Pada Function terdapat :
1. parameter / input
2. proses
3. hasil / return value
================= */

func printMyResult(sentence string) {
	fmt.Println(sentence)
}

// Multiple Return Value
func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

// Predefined Return Value
// mendefinisikan fariabel sejajar dengan function
func prCalculate(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return luas, keliling
}
