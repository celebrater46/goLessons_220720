// Interface
// If a struct has a method, it can use the struct as an interface
// Empty Interface can receive all types (L21)

package main

import "fmt"

// both types japanese and american are different type. So cannot put into Slice

// Empty
func show(t interface{}) {
	// Type assertions
	//_, isJapanese := t.(japanese) // value, bool := ...
	//if isJapanese {
	//	fmt.Println("I am Japanese")
	//} else {
	//	fmt.Println("I am not Japanese")
	//}

	// Type Switch
	switch t.(type) {
	case japanese:
		fmt.Println("I am Japanese")
	default:
		fmt.Println("I am not Japanese")
	}
}

type greeter interface {
	greet()
}

type japanese struct{}
type american struct{}

func (j japanese) greet() {
	fmt.Println("Konnichiwa!")
}
func (a american) greet() {
	fmt.Println("Hello!")
}

func main() {
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.greet() // Konnichiwa! Hello!
		//greeter.show() // Undefined
		show(greeter)
	}
	fmt.Println(greeters) // [{} {}]
}
