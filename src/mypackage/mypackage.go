package mypackage

import "fmt"

// CarPublic represents a publicly accessible car with a brand and year.
type CarPublic struct {
	Brand string
	Year  int
}

type CarPrivate struct {
	brand string
	year  int
}

// PrintMessage print a public message
func PrintMessage(text string) {
	fmt.Printf("Hello, I am %s\n", text)
}

func printMessage(text string) {
	fmt.Printf("Hello, I am %s\n", text)
}
