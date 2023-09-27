package main

import "fmt"

func isEven(number int) {
	if number%2 == 0 {
		fmt.Printf("%d es par \n", number)
	} else {
		fmt.Printf("%d es impar \n", number)
	}
}

func login(user, password string) bool {

	const userSaved string = "aldo.matus@outlook.com"
	const passwordSaved string = "Myp@ssword2023"

	if user == userSaved && password == passwordSaved {
		fmt.Println("Successful")
		return true
	} else if user == userSaved {
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
