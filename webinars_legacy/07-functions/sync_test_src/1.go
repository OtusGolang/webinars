package main

import (
	"fmt"
)

func f(v1, v2 int) {
}

func main() {
	var v1, v2 int
	{
		v1 = 1
		v2 := 2
		f(v1, v2)
	}
	fmt.Println(v1 + v2)
}
