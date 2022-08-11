package main

/*Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.*/
import (
	"fmt"
	"strconv"
)

func setIthBit(i int, bit int, number int64) int64 {
	if bit == 1 {
		return number | (1 << i)
	} else {
		return number &^ (1 << i)
	}
}

func main() {
	var i, bit int
	var number int64
	fmt.Println("Введите значение числа : ")
	fmt.Scan(&number)
	fmt.Println("Введите значение бита: ")
	fmt.Scan(&bit)
	fmt.Println("Введите значение вашего индекса: ")
	fmt.Scan(&i)
	fmt.Println("Before: \n", strconv.FormatInt(number, 2))
	number = setIthBit(i, bit, number)
	fmt.Println("Before: \n", strconv.FormatInt(number, 2))
}
