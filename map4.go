package main

import (
	"fmt"
)

func main() {
	m := map[int]int{
		1:  12,
		12: 5,
	}

	delete(m, 12)
	fmt.Println(m)
}
