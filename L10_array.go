// Array

package main

import "fmt"

func main() {
	var a [5]int // a[0] - a[4]
	a[2] = 3
	a[4] = 10
	fmt.Println(a[1]) // 0
	fmt.Println(a[2]) // 3
	fmt.Println(a[3]) // 0
	fmt.Println(a[4]) // 10
	// b := [3]int{1, 3, 5}
	b := [...]int{1, 3, 5}
	fmt.Println(b)
	fmt.Println(len(b)) // length
}
