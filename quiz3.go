package main

import (
	"errors"
	"fmt"
)

func Quiz3() {
	// membuat function sum
	// menjumlahkan semua elemen2 nya

	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)

	fmt.Println("Total : ", total)

	result, err := hitung(10, 2, "]")
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println("Hasil : ", result)
}

func sum(scores []int) (jum int) {
	for _, value := range scores {
		jum = jum + value
	}
	return jum
}

func hitung(angka1 int, angka2 int, simbol string) (int, error) {

	var hasil int
	var errorResult error

	switch simbol {
	case "+":
		hasil = angka1 + angka2
	case "-":
		hasil = angka1 - angka2
	case "*":
		hasil = angka1 * angka2
	case "/":
		hasil = angka1 / angka2
	default:
		errorResult = errors.New("Simbol tidak tersedia")
	}

	return hasil, errorResult
}
