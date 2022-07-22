// Operation
/*
num + - * / %
string +
bool AND(&&) OR(||) NOT(!)
*/

package main

import "fmt"

func main() {
	// INT
	var x int
	//x = 10 % 3
	x += 3 // x = x + 3
	x++    // Cannot write ++x
	fmt.Println(x)

	// STRING
	var s string
	s = "hello " + "world"
	fmt.Println(s) // hello world

	a := true
	b := false
	fmt.Println(a && b) // false
	fmt.Println(a || b) // true
	fmt.Println(!a)     // false
}
