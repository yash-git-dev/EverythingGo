package Controllers

import (
	"fmt"
	"regexp"
	"strings"
)

func RegexMain() {

	pattern := regexp.MustCompile(`\b\w+@\w+\.\w+\b`)

	match := pattern.MatchString("yash@gmail.com")

	if match {
		fmt.Println("matched")
	} else {
		fmt.Println("not matched")
	}

	text := "John's email is john.doe@example.com"
	submatches := pattern.FindStringSubmatch(text)

	for _, match := range submatches {
		fmt.Println(match)
	}
}

func ReverseStringWithSameWordLength(input string) string {
	// Split the original string into words
	words := strings.Fields(input)

	reversedWords := make([]string, len(words))
	for i := 0; i < len(words); i++ {
		reversedWords[i] = words[len(words)-1-i]
	}
	// Reverse the order of words
	for i, j := 0, len(reversedWords)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// // Reverse the characters in each word
	// for i, word := range reversedWords {
	// 	words[i] = reverseString(word)
	// }

	// Join the reversed words to form the reversed string
	reversedString := strings.Join(reversedWords, " ")

	return reversedString
}

func reverseString(s string) string {
	// Convert the string to a rune slice
	runes := []rune(s)

	// Reverse the order of characters
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
