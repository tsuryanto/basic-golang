package main

import "fmt"

func LatihanForControl() {
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Saya Sedang Belajar For Cotrol Go:", i)
	// }

	// //WHILE
	// i := 1
	// for i <= 100 {
	// 	fmt.Println()
	// 	i++
	// }

	title := "Golang the best language"

	// for index, letter := range title {
	// 	fmt.Println("index:", index, " letter : ", string(letter))
	// }

	// for _, letter := range title {
	// 	fmt.Println(" letter : ", string(letter))
	// }

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index : ", index, " letter : ", string(letter))
		}
	}

	for _, letter := range title {

		huruf := string(letter)

		switch huruf {
		case "a", "i", "u", "e", "o":
			fmt.Println(" letter : ", huruf)
		}

		if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
			fmt.Println(" letter : ", huruf)
		}

	}
}
