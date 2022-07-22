// for
// Golang has only "for". No "while".

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		//if i == 3 { break } // Stop forever
		//if i == 5 { continue } // Stop once
		fmt.Println(i) // 0 - 9
	}

	// Can use like "while"
	j := 11
	for j < 20 {
		fmt.Println(j) // 11 - 19
		j++
	}

	// Endless loop
	k := 21
	for {
		fmt.Println(k) // 21- 30
		k++
		if k > 30 {
			break
		}
	}
}
