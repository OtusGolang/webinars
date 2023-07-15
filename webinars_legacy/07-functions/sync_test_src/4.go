package main

import (
	"fmt"
)

func main() {
	s := 10
	defer fmt.Println(s)
	s++
}
