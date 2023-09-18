package main

import "fmt"

func main() {
	nums := [3]int{1, 2, 3}

	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])

	nums[1] = 33

	fmt.Println(nums)
}
