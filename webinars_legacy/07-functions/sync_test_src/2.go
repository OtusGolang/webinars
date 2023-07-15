package main2

import "fmt"

func f(x int) {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}

func main() {
	var x = 200
	f(x)
}
