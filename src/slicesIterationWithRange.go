package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {

	var textReverse string
	lowercaseText := strings.ToLower(text)
	textWithoutSpaces := strings.ReplaceAll(lowercaseText, " ", "")

	for i := len(textWithoutSpaces) - 1; i >= 0; i-- {
		textReverse += string(textWithoutSpaces[i])
	}

	if textReverse == textWithoutSpaces {
		fmt.Println("It is Palindrome")
	} else {
		fmt.Println("It is not a Palindrome")
	}

}

func main() {
	slice := []string{"How", "is", "it", "going"}

	for _, value := range slice {
		fmt.Println(value)
	}

	isPalindrome("saa")
}
