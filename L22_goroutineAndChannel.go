// Goroutine
// Channel (L23)
// Goroutine gives a parallel execution
// Channel is like a pipe to carry data

package main

import (
	"fmt"
	"time"
)

//func task1() {
func task1(result chan string) { // Channel
	time.Sleep(time.Second * 2) // Waiting 2 sec
	//fmt.Println("task1 finished!")
	result <- "task1 result"
}

func task2() {
	fmt.Println("task2 finished!")
}

func main() {
	// task1() -> task2(), task2() waited until task1() finished
	//task1()
	//task2()

	// go == goroutine, put it before func()
	//go task1()
	//go task2()

	result := make(chan string)

	go task1(result)
	go task2()

	fmt.Println(<-result) // task2 finished!, task1 result

	// go task1() and go task2() did not execute unless adding this time.Sleep()
	time.Sleep(time.Second * 3)
}
