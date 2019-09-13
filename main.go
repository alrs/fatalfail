package main

import (
	"fmt"
	"time"
)

func main() {
	a := 2
	b := 2
	fmt.Printf("%d + %d = %d\n", a, b, addTwo(a, b))
}

func addTwo(a int, b int) int {
	time.Sleep(1 * time.Second)
	return a + b
}
