


func main() {
	ch:=make(chan int)
	go func() {
		fmt.Print(7)
		ch <- 5
	}()
	fmt.Print(<-ch)
}


func main() {
	ch := make(chan int, 3)
	ch <- 19
	ch <- 27
	ch <- 53
	close(ch)

	for v := range ch {
		fmt.Print(v)
	}
}


func main() {
	ch1 := make(chan int, 10)
	ch1 <- 1

	ch2 := make(chan int, 10)
	ch2 <- 2

	select {
	case ch1 <- 3:
		fmt.Println("1")
	case <-ch2:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}
}



func main() {
	ch := make(chan int, 10)
	ch <- 3
	close(ch)

	v, ok := <-ch
	fmt.Print(v)

	v, ok = <-ch
	fmt.Print(v)

	if ok {
		fmt.Print("!")
	}
}

