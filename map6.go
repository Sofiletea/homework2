package main

import "fmt"

func main() {
	m := map[int]bool{}
	for i := 0; i < 5; i++ {
		m[i] = ((i % 2) == 0)
	}
	for k, v := range m {
		fmt.Printf("key: %d, value: %t\n", k, v)
	}
}
