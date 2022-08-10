package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func CloseChan(ch <-chan int, jr *sync.WaitGroup) {
	for i := range ch {
		fmt.Print(i, " ")

	}
	fmt.Println("Нет горутинам")
	jr.Done()
	return
}

func Chan(exit chan interface{}, jr *sync.WaitGroup) {
	for {
		select {
		case <-exit:
			fmt.Println("Нет горутинам")
			jr.Done()
			return
		default:
			fmt.Println("Горутины все-таки живы")
			time.Sleep(1 * time.Second)
		}
	}
}

func Context(bn context.Context, jr *sync.WaitGroup) {
	for {
		select {
		case <-bn.Done():
			fmt.Println("Горутинам нет")
			jr.Done()
			return
		default:
			fmt.Println("Горутины все-таки живы")
			time.Sleep(1 * time.Second)
		}
	}

}

func main() {
	var jr sync.WaitGroup
	jr.Add(1)
	ch := make(chan int)
	go CloseChan(ch, &jr)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	jr.Wait()
	jr.Add(1)
	exit := make(chan interface{})

	go Chan(exit, &jr)
	time.Sleep(1 * time.Second)
	exit <- struct{}{}
	jr.Wait()
	bn, cancel := context.WithCancel(context.Background())
	jr.Add(1)
	go Context(bn, &jr)
	time.Sleep(5 * time.Second)
	cancel()
	jr.Wait()
}
