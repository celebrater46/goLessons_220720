// Switch

package main

import "fmt"

func main() {
	signal := "red"
	switch signal {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Caution")
	case "green", "blue":
		fmt.Println("Go")
	default:
		fmt.Println("wrong signal")
	}

	score := 82
	switch {
	case score > 80:
		fmt.Println("Great!")
	default:
		fmt.Println("so so ...")
	}

}
