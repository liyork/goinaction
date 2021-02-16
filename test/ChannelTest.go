package test

import (
	"fmt"
	"math/rand"
	"time"
)

func testChannel() {

	//无缓冲整形通道
	unbuffered := make(chan int)
	fmt.Print(unbuffered)

	//有缓冲字符串通道
	buffered := make(chan string, 10)

	//发送字符串到通道
	buffered <- "abfcf"

	//从通道接收字符串
	value := <-buffered
	fmt.Print(value)
}

func Relay() {

	baton := make(chan int)
	wg.Add(1)

	go Runner(baton)

	//通道写入
	baton <- 1

	//等待直到最后一个赛手完成
	wg.Wait()

	//runner := <-baton
	//fmt.Printf("the finale baton:%d", runner)
}

//初始化
func init() {
	rand.Seed(time.Now().UnixNano())
}

func PingPon() {
	//无缓冲通道
	court := make(chan int)
	wg.Add(2)

	go player("aaa", court)
	go player("bbb", court)

	//启动
	court <- 1

	//等待都结束
	wg.Wait()
}

func player(name string, court chan int) {

	defer wg.Done()

	for {
		ball, ok := <-court //阻塞等待直到有值
		if !ok {            //通道关闭
			fmt.Printf("Player %s Won \n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 { //若非被13整除则失败，关闭通道
			fmt.Printf("Player %s Missed \n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d \n", name, ball)
		ball++

		//放入通道
		court <- ball
	}
}

func Runner(baton chan int) {
	var newRunner int

	//无缓冲通道，阻塞直到有数据
	runner := <-baton

	fmt.Printf("Runner %d Running With Baton\n", runner)

	//单开线程执行
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The line\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(100 * time.Microsecond)

	//最后一人完成
	if runner == 4 {
		fmt.Printf("Runner %d Finished,Race Over\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	//新值赋值给通道
	baton <- newRunner
}
