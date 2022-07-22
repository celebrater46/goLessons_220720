// Map
// It has Key and Value like JS Object

package main

import "fmt"

func main() {
	//m := make(map[string]int) // Key is string type, Value is int type
	//m["Hizuru"] = 170
	//m["Seiko"] = 165
	m := map[string]int{"Hizuru": 170, "Seiko": 165} // Same above
	fmt.Println(m)                                   // map[Hizuru:170 Seiko:165]
	fmt.Println(len(m))                              // 2
	delete(m, "Seiko")
	fmt.Println(m) // map[Hizuru:170]
	cm, exists := m["Hizuru"]
	fmt.Println(cm)     // 170
	fmt.Println(exists) // true
}
