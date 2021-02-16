package test

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter4 int64
	wg4      sync.WaitGroup
	mutex    sync.Mutex
)

func RaceConditionLockTest() {
	wg4.Add(2)

	go incCounterLock(1)
	go incCounterLock(2)

	fmt.Println("Waiting To Finish")
	wg4.Wait()

	fmt.Println("Final Counter4:", counter4)
}

//同一时刻只有一个 goroutine 可以进入临界区。之后，直到调用 Unlock()函数之后，其他 goroutine 才能进入临界区。
// 若强制将当前 goroutine 退出当前线程后，调度器会再次分配这个 goroutine 继续运行。因为它获取锁了
func incCounterLock(id int) {
	defer wg4.Done()

	for count := 0; count < 200; count++ {
		mutex.Lock()
		{
			value := counter4
			runtime.Gosched()
			value++
			counter4 = value
		}
		mutex.Unlock()
	}
}
