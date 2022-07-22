// range
// It is like "foreach"

package main

import "fmt"

func main() {
	s := []int{2, 3, 8}
	for key, value := range s {
		fmt.Println(key, value) // 0 2, 1 3, 2 8
	}

	s2 := []int{1, 2, 7}
	for _, value := range s2 { // "_" ignores "key"
		fmt.Println(value) // 1, 2, 7
	}

	m := map[string]int{"Reina": 53, "Sally": 34}
	for key, value := range m {
		fmt.Println(key, value) // Reina 53, Sally 34

	}
}
