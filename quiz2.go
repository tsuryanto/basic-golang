package main

import "fmt"

func Quiz2() {

	/* SOAL
	1. Hitung Rata-rata. Tampilkan nilai koma
	2. Jika score score >= 90, tambahkan ke var goodScores
	*/

	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var goodScores []int
	var total int

	for _, score := range scores {
		total = total + score
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}

	fmt.Println("Total : ", total)
	fmt.Println("Rata-rata : ", float64(float64(total)/float64(len(scores))))
	fmt.Println("Good Scores : ", goodScores)
}
