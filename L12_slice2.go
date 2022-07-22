// Slice2
// make(), append(), copy()

package main

import "fmt"

func main() {
	// Make Slice
	//s := make([]int, 3) // [0 0 0]
	s := []int{1, 3, 5} // If [] is empty, "s" will become Slice
	fmt.Println(s)      // 1 3 5

	// appendJ() means adding some value at the end of the array
	s = append(s, 8, 2, 10)
	fmt.Println(s) // 1 3 5 8 2 10

	// copy() means
	t := make([]int, len(s)) // Slice to copy
	//copy(t, s) // Copy s to t
	n := copy(t, s) // n == len(t)
	fmt.Println(s)  // 1 3 5 8 2 10
	fmt.Println(t)  // 1 3 5 8 2 10
	fmt.Println(n)  // 6
}
