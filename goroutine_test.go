package gogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello, World!")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()

	fmt.Println("TestCreateGoroutine: Hello, World!")

	time.Sleep(1 * time.Second)
}

func DisplayNumbers(number int) {
	fmt.Println("Displaying Number: ", number)
}

func TestManyGorountine(t *testing.T) {
	for i := 0; i <= 100000; i++ {
		go DisplayNumbers(i)
	}

	time.Sleep(10 * time.Second)
}
