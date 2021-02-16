package test

import (
	"fmt"
	"runtime"
	"sync"
)

//可以使用go build -race检查

var (
	counter int
	wg      sync.WaitGroup
)

func TestRaceCondition() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Final Counter:", counter) //结果不一定是400
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 200; count++ {
		//临时存储
		value := counter
		//让当前goroutine从线程退出，并放回到队列
		runtime.Gosched()
		value++
		counter = value
	}
}
