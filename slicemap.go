package main

import "fmt"

func LatihanSliceMap() {
	students := []map[string]string{
		{"name": "Taufiq", "score": "A"},
		{"name": "Saefudin", "score": "B"},
		{"name": "Dandy", "score": "C"},
	}

	for _, student := range students {
		fmt.Println(student["name"], " = ", student["score"])
	}
}
