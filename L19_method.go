// method

package main

import "fmt"

type user2 struct {
	name  string
	score int
}

// Method is written at out of the struct
func (u user2) show() {
	fmt.Printf("name: %s, score: %d\n", u.name, u.score)
}

//func (u user2) hit() { // This function cannot overwrite fields
func (u *user2) hit() { // Paste the Link (The function can overwrite)
	u.score++
}

func main() {
	u := user2{name: "Muramasa", score: 320}
	u.hit()
	u.show()
}
