package main

import "fmt"

func Partition(mass []int, low, high int) int {
	pivot := mass[high]
	m := low
	for j := low; j < high; j++ {
		if mass[j] < pivot {
			mass[m], mass[j] = mass[j], mass[m]
			m++
		}
	}
	mass[m], mass[high] = mass[high], mass[m]
	return m
}

func quickSort(mass []int, low, high int) {
	if low < high {
		p := Partition(mass, low, high)
		quickSort(mass, low, p-1)
		quickSort(mass, p+1, high)
	}
}

func main() {
	mass := []int{13, 10, 8, 12, 6, 4, 15, 11, 5, 3, 7, 2, 9, 1}
	fmt.Println("Before sorting: %v\n", mass)
	quickSort(mass, 0, len(mass)-1)
	fmt.Println("After sorting: %v\n", mass)
}
