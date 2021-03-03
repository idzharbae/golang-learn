package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	
	go func(ctx context.Context) {
		tick := time.NewTicker(time.Second)
		for {
			select {
			case <-tick.C:
				fmt.Println("TICK")
			case <-ctx.Done():
				fmt.Println("TOCK")
				return
			}
		}
	}(ctx)
	
	time.Sleep(time.Second*10)
	
	cancelFunc()
	
	fmt.Println("Hello, playground")
	
	time.Sleep(time.Second)
}
