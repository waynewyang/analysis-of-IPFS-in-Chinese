package main

import (
	"context"
	"fmt"
	"time"
)

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("end and exit!")
			return
		default:
			fmt.Println(name, "running")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go watch(ctx, "proc1")
	go watch(ctx, "proc2")
	go watch(ctx, "proc3")
	go watch(ctx, "proc4")
	go watch(ctx, "proc5")

	time.Sleep(10 * time.Second)
	fmt.Println("end request")
	cancel()

	time.Sleep(5 * time.Second)
}
