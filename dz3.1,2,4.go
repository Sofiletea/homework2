package main

import (
	"fmt"
)

const glo int = 10

func main() {
	const loc int = 3
	var result int = glo + loc
	var result1 int = glo - loc
	var result2 int = glo * loc
	var result3 int = glo / loc

	fmt.Println(glo, loc, result, result1, result2, result3)
}
