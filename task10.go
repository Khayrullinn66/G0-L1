package main

import "fmt"

func Mapa(mass []float64) map[int][]float64 {
	data := make(map[int][]float64, 0)
	for _, i := range mass {
		data[int(i/10)*10] = append(data[int(i/10)*10], i)
	}
	return data
}

func main() {
	mass := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println("%-#v\n", Mapa(mass))
}
