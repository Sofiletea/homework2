package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	sl := []string{"apple", "banana", "cherry"}
	sl[0] = "mango"
	fmt.Println(slice, sl)
}
