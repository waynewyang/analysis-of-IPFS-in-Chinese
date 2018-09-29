package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case i := <-stop:
				fmt.Println(i)
				fmt.Println("program end!")
				return

			default:
				fmt.Println("running......")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("end request")
	stop <- false

	time.Sleep(5 * time.Second)
}
