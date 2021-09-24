package main

import "fmt"

func LatihanMap() {
	// // CARA 1
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["ruby"] = 90
	// myMap["javascript"] = 80
	// myMap["go"] = 100

	// fmt.Println(myMap)

	// CARA 2
	myMap := map[string]string{
		"ruby":       "is beautiful",
		"go":         "is super fast",
		"javascript": "hype",
	}
	fmt.Println(myMap["go"])

	for key, value := range myMap {
		fmt.Println("Key : ", key, "Value : ", value)
	}

	fmt.Println("==========")

	// Delete Map Value
	delete(myMap, "ruby")
	for key, value := range myMap {
		fmt.Println("Key : ", key, "Value : ", value)
	}

	fmt.Println("==========")

	// Cek apakah suatu map ada value nya atau tidak
	value, apakahAda := myMap["python"]
	fmt.Println(apakahAda)
	fmt.Println(value)
}
