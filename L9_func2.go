// Func2
// Func in var
// Func In Func
// Quick No Named Func like JavaScript

package main

import "fmt"

//func swap(a, b int) (int, int) {
//	return b, a // Can return VALUES (Not only one value)
//}

func main() {
	//fmt.Println(swap(5,2))

	// Func in var
	//f := func(a, b int) (int, int) {
	//	return b, a
	//}
	//fmt.Println(f(2, 3))

	// Quick No Named Func
	func(msg string) {
		fmt.Println(msg)
	}("Hideru")
}
