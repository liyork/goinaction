package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//原子访问int64变量

var (
	shutdown int64
	wg3      sync.WaitGroup
)

func AtomicLoadStoreTest() {
	wg2.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(2 * time.Second)

	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	wg3.Wait()
}

func doWork(name string) {
	defer wg3.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Microsecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down \n", name)
			break
		}
	}
}
