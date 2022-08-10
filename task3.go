package main

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2* 2 + 3 * 3 + 4 * 4 …)
с использованием конкурентных вычислений.
*/
import "fmt"

func main() {
	mass := [...]int{2, 4, 6, 8, 10}
	do := make(chan int, len(mass))
	defer close(do)
	// запуск горутин, которые считают произведние и записывают в канал
	for _, i := range mass {
		go func(i int) {
			do <- i * i
		}(i)
	}
	sum := 0
	// считают данные из канала и их сумму
	for i := 0; i < len(mass); i++ {
		sum += <-do
	}
	fmt.Print(sum, "\n")
}