package main

import "fmt"

func gen(v1, v2 int) func() int {
	secret := 7
	return func() int {
		secret--
		return v1 + v2 + secret
	}
}

func main() {
	v1 := 1
	v2 := 2
	f := gen(v1, v2)
	f()
	fmt.Println(f())
}
