package otus_test

import (
	"fmt"
	"testing"
)

func TestAtomicStoreLoad(t *testing.T) {
	ch := make(chan struct{})

	var v int32
	go func() {
		v = 1
		close(ch)
	}()
	fmt.Println(v)

	<-ch
}

func TestAtomicAdd(t *testing.T) {
	ch := make(chan struct{})

	var v int32
	go func() {
		v++
		close(ch)
	}()
	fmt.Println(v)

	<-ch
}

func TestAtomicCAS(t *testing.T) {
	ch := make(chan struct{})

	var v int32
	go func() {
		v = 1
		close(ch)
	}()

	for v != 1 {
	}
	v = 2
	fmt.Println(v)

	<-ch
}
