package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	context.TODO()
	ctx = context.WithValue(ctx, "a", "b")
	ctx = context.WithValue(ctx, "c", "d")
	ctx, cancelFunc := context.WithCancel(ctx)
	//ctx2, _ := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	go func() {
		time.Sleep(3 * time.Second)
		cancelFunc()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Done")
	}

	//ctx3, cancelFunc := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	//ctx4, cancelFuun := context.WithTimeout(ctx, time.Minute)
}
