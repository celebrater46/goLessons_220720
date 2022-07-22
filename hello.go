// Must write "package hoge" in Go file
// Must name "package main" one of files

//package test
package main

// For input and output
import (
	"fmt"
	"math/rand"
	"time"
)

/*
Compile -> go build hello.go
Compile and Run -> go run hello.go
*/

func main() {
	// Println means adding \n after the words
	fmt.Println("Hello World.")

	//var name string = "Hizuru"
	//var name = "Hizuru"
	name := "Hizuru"
	fmt.Printf("Hello %v\n", name)
	fmt.Printf("Hello %v again!\n", name)

	// Receive User's input
	//var guess int
	//var answer int = 7
	//rand.Seed(time.Now().UnixNano()) // If this method is nothing, the answer is same number every time.
	//var answer int = rand.Intn(10)   // 0-9
	//fmt.Print("Type a number you guess: ")
	//fmt.Scanf("%v", &guess)
	//fmt.Printf("You pressed %v\n", guess)

	//if answer == guess {
	//	fmt.Println("Correct.")
	//} else if answer > guess {
	//	fmt.Println("The answer is higher.")
	//} else {
	//	fmt.Println("The answer is lower.")
	//}

	//fmt.Printf("The answer is %v\n", answer)

	rand.Seed(time.Now().UnixNano()) // If this method is nothing, the answer is same number every time.
	var answer2 int = rand.Intn(10)  // 0-9
	var count int = 0

	for {
		var guess2 int

		fmt.Print("Type a number you guess2: ")
		fmt.Scanf("%v", &guess2)
		fmt.Printf("You pressed %v\n", guess2)
		//count = count + 1
		//count += 1
		count++

		if answer2 == guess2 {
			//fmt.Println("Correct.")
			fmt.Printf("Correct. It took %v guesses.\n", count)
			break
		} else if answer2 > guess2 {
			fmt.Println("The answer is higher.")
		} else {
			fmt.Println("The answer is lower.")
		}

		// For Test
		//fmt.Printf("The answer is %v\n", answer2)
	}
}
