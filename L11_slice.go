// Slice
// To save Memory

package main

import "fmt"

func main() {
	a := [5]int{2, 10, 8, 15, 4}
	s := a[2:4]
	s2 := a[2:]
	s3 := a[:4]
	s4 := a[:]

	fmt.Println(a)
	fmt.Println(s)  // [8, 15]
	fmt.Println(s2) // [8, 15, 4]
	fmt.Println(s3) // [2, 10, 8, 15]
	fmt.Println(s4) // [2, 10, 8, 15, 4]

	// Slice is Link. Not being copied the value
	s[1] = 95
	fmt.Println(a)      // [2, 10, 8, 95, 4]
	fmt.Println(s)      // [8, 95]
	fmt.Println(len(s)) // 2
	fmt.Println(cap(s)) // 3
}
