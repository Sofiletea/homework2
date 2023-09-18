package main

import (
	"fmt"
)

func main() {
	const loc int = 3
	//loc = 12
	//нельзя менять - бедет ошибка
	const m float64 = 22.3
	const t string = "test"
	const b bool = true

	fmt.Println(loc, m, t, b)
}
