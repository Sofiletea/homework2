package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	subSlice := slice[1:3]
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	fmt.Println("Скопированный слайс:", slice2)
	fmt.Println(subSlice)
}
