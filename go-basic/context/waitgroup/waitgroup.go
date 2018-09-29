package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1 finished")
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2 finished")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("all finished")
}
