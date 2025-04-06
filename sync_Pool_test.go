package gogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New Data" // Default value when no data is available in the pool
		},
	}

	pool.Put("Hello")
	pool.Put("World")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()

			fmt.Println("Data: ", data)
			// Simulate some processing
			time.Sleep(2 * time.Millisecond)
			pool.Put(data)
		}()
	}

	// Wait for goroutines to finish
	time.Sleep(1 * time.Second)
}
