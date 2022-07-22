package main

import "fmt"

// Basic Data Type
// string  "Hello"
// int  53
// float64  10.2
// bool
// nil

// Empty Data
var s string // ""
var a int    // 0
var f bool   // false

func main() {
	a := 10
	b := 12.3
	c := "Hoge"
	var d bool
	fmt.Printf("a: %d, b: %f, c: %s, d: %t\n", a, b, c, d)
}
