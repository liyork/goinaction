package fmt

import (
	"fmt"
)

//未公开
type alertCounter int
//公开
type AlertCounter1 int

func Println(a ...interface{}) () {
	fmt.Print(a)
}

//获取未公开属性
func GetAlertCounter() alertCounter {
	return alertCounter(1)
}

//将工厂函数命名为 New 是 Go 语言的一个习惯
func New(value int) alertCounter{
	return alertCounter(value)
}
