package main

import "fmt"

func LatihanArray() {
	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "Javascript"
	// languages[3] = "C#"
	// languages[4] = "Python"

	// languages := [5]string{"Go", "Ruby", "Javascript", "C#", "Python"}

	// isi pakai ... jika ingin jumlah index nya otomatis dibuatkan tergantung banyak datanya
	languages := [...]string{
		"Go",
		"Ruby",
		"Javascript",
		"C#",
		"Kotlin",
		"Python", // harus pakai koma , jika vertikal
	}

	fmt.Println(languages)

	// HITUNG LENGTH Array
	length := len(languages)
	fmt.Println(length)

	for index, lang := range languages {
		fmt.Println("Index: ", index, " Language: ", lang)
	}
}
