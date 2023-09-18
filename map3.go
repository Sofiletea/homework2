package main

import "fmt"

func main() {
	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println(people["Alice"]) // 8
	fmt.Println(people["Bob"])   // 2
	people["Bob"] = 32
	fmt.Println(people["Bob"])
}
