// Const

package main

import "fmt"

func main() {
	//const name = "Hideru"
	//name = "Miyon"
	//fmt.Println(name) // # command-line-arguments ./L5_const.go:9:7: cannot assign to name (declared const)

	// iota means to fill an array 0, 1, 2...
	const (
		sun = iota // 0
		mon        // 1
		tue        // 2
	)
	fmt.Println(sun, mon, tue)
}
