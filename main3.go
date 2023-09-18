package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice)) // Выводит 5
	fmt.Println(cap(slice)) // Выводит 5
}
