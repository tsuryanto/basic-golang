package main

import "fmt"

func LatihanSlice() {

	// FORMAT SLICE 1
	var gamingConsoles []string
	gamingConsoles = append(gamingConsoles, "Playstation 4")
	gamingConsoles = append(gamingConsoles, "Nintendo Switch")
	gamingConsoles = append(gamingConsoles, "Xbox One")

	// // FORMAT SLICE 2
	// gamingConsoles := []string{"Playstation 4", "Nintendo Switch", "Xbox One"}

	// fmt.Println(gamingConsoles)

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}
}
