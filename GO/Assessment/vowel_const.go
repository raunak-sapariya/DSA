package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Enter number of characters: ")

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error: Please enter a valid integer.")
		return
	}

	var vowels []string
	var consonants []string

	for i := 0; i < n; i++ {
		var char string

		fmt.Printf("Enter character %d: ", i+1)

		_, err = fmt.Scan(&char)
		if err != nil {
			fmt.Println("Error: Invalid input.")
			i--
			continue
		}

		if len(char) != 1 {
			fmt.Println("Error: Please enter only one character.")
			i--
			continue
		}

		c := char[0]
		if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) {
			fmt.Println("Error: Please enter an alphabet only.")
			i--
			continue
		}

		switch char {
		case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
			vowels = append(vowels, char)
		default:
			consonants = append(consonants, char)
		}
	}

	fmt.Println("\nVowels:", vowels)
	fmt.Println("Consonants:", consonants)
}