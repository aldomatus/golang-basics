package main

import "fmt"

func isEven(number int) {
	if number%2 == 0 {
		fmt.Printf("%d es par \n", number)
	} else {
		fmt.Printf("%d es impar \n", number)
	}
}

func login(username, password string) bool {

	const usernameSaved string = "aldo.matus@outlook.com"
	const passwordSaved string = "Myp@ssword2023"

	if username == usernameSaved && password == passwordSaved {
		fmt.Println("Successful")
		return true
	} else if username == usernameSaved {
		fmt.Println("Invalid password")
		return false
	} else {
		fmt.Println("Invalid username")
		return false
	}

}

func main() {

	// Test if a number is even
	const number int = 13
	isEven(number)

	// Test login system
	const username string = "aldo.matus@outlook.com"
	const password string = "Myp@ssword2022"
	login(username, password)

}
