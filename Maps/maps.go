package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	words := map[string]int64{}

	s := strings.Fields(strings.ToLower(text))

	/*
		for _, value := range s {
			mapValue, ok := words[value]
			if !ok {
				words[value] = 1
			} else {
				words[value] = mapValue + 1
			}
		}
	*/

	for _, value := range s {
		words[value]++
	}

	fmt.Println(words)
}
