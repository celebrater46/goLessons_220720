// if
//> >= < <= == != && || !

package main

import "fmt"

func main() {
	score := 83

	//if score := 43; score > 80 { // If declared a var in "if" Scope, the var is available only in the Scope
	if score > 80 {
		fmt.Println("Great!")
	} else if score > 60 {
		fmt.Println("Good.")
	} else {
		fmt.Println("so so ...")
	}
}
