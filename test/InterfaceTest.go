package test

import (
	"fmt"
)

//定义了通知类行为的接口
type notifier interface {
	notify()
}

type user2 struct {
	name  string
	email string
}

func (u *user2) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin2 struct{
	name string
	email string
}

//使用指针接收者实现notifier接口
func (a *admin2) notify(){
	fmt.Printf("Sending admin email to %s<%s>\n",a.name,a.email)
}

//多态
func testInterface(){
	u := user2{"bill","bill@emai.com"}

	//sendNotification(u)//编译出错，因为值中不包含指针的方法集
	sendNotification(&u)

	lisa := admin2{"lisa","bill@emai.com"}
	sendNotification(&lisa)
}

//接收实现了notifier接口的值
func sendNotification(n notifier) {
	n.notify()
}
