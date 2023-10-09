package main

import "fmt"

func main() {
	// Initialize an empty map with keys of type string and values of type int.
	m := make(map[string]int)
	m["Aldo"] = 26
	m["Melina"] = 28

	fmt.Println("m: ", m)

	// Loop through the map and print each key-value pair.
	// 'k' holds the key and 'v' holds the value for each iteration.
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Retrieve a value from the map.
	// 'value' holds the value, and 'ok' holds a boolean indicating if the key was found.
	value, ok := m["Aldo"]
	fmt.Println(value, ok)

}
