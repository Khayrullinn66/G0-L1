package main

import "fmt"

func main() {
	x := 5
	y := 10
	fmt.Println("Before : x = %d, y = %d\n", x, y)
	x, y = y, x
	fmt.Println("After change 1: x + %d, y = %d\n", x, y)

	x = x + y
	y = x - y
	x = x - y
	fmt.Println("After change 2: x = %d, y = %d\n", x, y)
}
