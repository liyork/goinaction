package test

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//可以使用go build -race检查

var (
	counter2 int64
	wg2      sync.WaitGroup
)

func RaceConditionAtomicTest() {
	wg2.Add(2)

	go incCounterAtomic(1)
	go incCounterAtomic(2)

	fmt.Println("Waiting To Finish")
	wg2.Wait()

	fmt.Println("Final Counter2:", counter2) //结果不一定是400
}

func incCounterAtomic(id int) {
	defer wg2.Done()

	for count := 0; count < 200; count++ {
		//原子增加
		atomic.AddInt64(&counter2, 1)
		//让当前goroutine从线程退出，并放回到队列
		runtime.Gosched()
	}
}
