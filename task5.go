package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Print("Введите значение : ")
	fmt.Scan(&n)
	ch := make(chan int)

	go func() {
		i := 0
		for {
			ch <- i
			i++
		}
	}()

	timer := time.After(time.Duration(n) * time.Second)

	for {
		select {
		case <-timer:
			return
		default:
			fmt.Println("Получить сообщение: %d\n", <-ch)
		}
	}
}
