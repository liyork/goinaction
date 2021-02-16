package test

import (
	"fmt"
	"runtime"
	"sync"
)

func TestGorouting() {
	//设定使用cpu数，供调度器使用，单cpu会在其上有多个gorouting并发执行，执行分片。
	runtime.GOMAXPROCS(1)//2
	//runtime.GOMAXPROCS(runtime.NumCPU())//每个逻辑处理器分配一个物理处理器

	//用来等待程序完成
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Gorutines")

	//匿名函数，使用goroutine执行
	go func() {
		//函数退出时执行
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println()
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println()
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
