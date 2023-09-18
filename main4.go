package main

import "fmt"

func main() {
	slice := make([]int, 5)
	slice = append(slice, 6)
	fmt.Println(slice)
}
