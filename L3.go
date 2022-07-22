package main

import "fmt"

// If first charactor of var's Name is a capital, it can watch the var from other package.

func main() {
	//var msg string
	//msg = "Hello World!"
	//var msg = "Hello World"
	msg := "Hello World!!!!!!"
	fmt.Println(msg)

	//var a, b int
	//a, b = 15, 10
	a, b := 15, 10

	var (
		c int
		d string
	)
	c = 10
	d = "hoge"
	fmt.Printf("a: %v, b: %v, c: %v, d: %v\n", a, b, c, d)
}
