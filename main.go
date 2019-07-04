// Given a word, compute the scrabble score for that word.
package main

import (
	"fmt"
	"os"
	"unicode"
)

var letterValues = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1,
	'T': 1, 'D': 2, 'G': 2, 'B': 3, 'C': 3, 'M': 3, 'P': 3, 'F': 4, 'H': 4,
	'V': 4, 'W': 4, 'Y': 4, 'K': 5, 'J': 8, 'X': 8, 'Q': 10, 'Z': 11,
}

func Score(input string) int {

	var totalScore int

	for _, w := range input {
		totalScore += letterValues[unicode.ToUpper(w)]
	}

	return totalScore
}

func main() {

	if len(os.Args) > 1 {
		inputWord := os.Args[1]
		fmt.Println(Score(inputWord))
	} else {
		fmt.Println("Error: specify a word")
	}

}
