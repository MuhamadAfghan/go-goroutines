package gogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	Mutex   sync.Mutex
	Balance int
	Name    string
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(from *UserBalance, to *UserBalance, amount int) {
	from.Lock()
	defer from.Unlock()

	to.Lock()
	defer to.Unlock()

	from.Change(-amount)
	to.Change(amount)
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{Name: "User 1", Balance: 100}
	user2 := UserBalance{Name: "User 2", Balance: 100}

	for i := 0; i < 10; i++ {
		go Transfer(&user1, &user2, 10)
		go Transfer(&user2, &user1, 10)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("User 1 Balance:", user1.Balance)
	fmt.Println("User 2 Balance:", user2.Balance)
}

// The above code demonstrates a potential deadlock situation where two goroutines are trying to lock the same resources in different order.
// This can lead to a situation where both goroutines are waiting for each other to release the locks, causing a deadlock.
