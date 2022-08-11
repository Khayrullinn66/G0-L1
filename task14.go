package main

import "fmt"

func findType(variable interface{}) {
	switch variable.(type) {
	case int:
		fmt.Println("Type int: %d\n", variable)
	case string:
		fmt.Println("Type string: %s\n", variable)
	case chan int:
		fmt.Println("Type chan int: %v\n", variable)
	default:
		fmt.Println("Unexpected type: %v\n", variable)
	}
}
func main() {
	findType(interface{}(45))
	findType(interface{}("mama"))
	findType(interface{}(true))
	findType(interface{}(make(chan int)))
	findType(interface{}(8.9))

}
