package main

/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.*/
import "fmt"

func main() {
	mass := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := range mass {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for val := range ch1 {
			ch2 <- val * 2
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}
}
