package main

/*Реализовать пересечение двух неупорядоченных множеств.*/
import "fmt"

func SetIntersection(arr1, arr2 []int) []int {
	arr3 := []int{}
	data := make(map[int]bool, len(arr1))

	for _, i := range arr1 {
		data[i] = true
	}
	for _, i := range arr2 {
		if data[i] {
			arr3 = append(arr3, i)
		}
	}
	return arr3

}
func main() {
	arr1 := []int{5, 9, 2, 4, 1, 6, 3, 8, 7}
	arr2 := []int{4, 1, 3, 2}
	fmt.Println(SetIntersection(arr1, arr2))
}
