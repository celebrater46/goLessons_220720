// Pointer
// It means to memorize the address of the Memory
// Cannot operate

package main

import "fmt"

func main() {
	a := 5
	var pa *int
	pa = &a // &a means the address of "a"
	// Write "*pa" to call the data in var pa
	fmt.Println(pa)  // 0xc000016048
	fmt.Println(*pa) // 5
}
