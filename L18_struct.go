// struct

package main

import "fmt"

type user struct {
	// Data that is included in struct called "Field"
	name string
	age  int
}

func main() {
	u := new(user)
	//(*u).name = "Tiki"
	u.name = "Tiki" // Same above
	u.age = 18
	fmt.Println(u) // &{Tiki 18} // Returned "pointer"

	//u2 := user{name: "Kirara", age: 28}
	u2 := user{"Kirara", 28}
	fmt.Println(u2) // {Kirara 28} // Returned "data" only. "pointer" was nothing
}
