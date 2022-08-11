package main

import "fmt"

func MakeSet(mass []string) []string {
	data := make(map[string]bool)
	for _, i := range mass {
		data[i] = true
	}
	deq := make([]string, 0, len(data))
	for key := range data {
		deq = append(deq, key)
	}
	return deq
}

func main() {
	mass := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(MakeSet(mass))
}
