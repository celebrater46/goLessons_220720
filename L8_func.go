// Func

package main

import "fmt"

/*
//func hi (name string) {
func hi (name string) string { // 2nd string means the type of the value to return
	//fmt.Println("Hi!" + name)
	msg := "hi!" + name
	return msg
}
*/

func hi(name string) (msg string) {
	// fmt.Println("hi!" + name)
	msg = "hi!" + name
	return
}

func main() {
	//hi("Hizuru")
	fmt.Println(hi("Hizuru"))
}
