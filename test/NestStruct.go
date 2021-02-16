package test

import (
	"fmt"
)

type notifier3 interface {
	notify()
}

type user3 struct {
	name  string
	email string
}

//可以通过user3类型值的指针调用的方法
func (u *user3) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin3 struct {//外部类型
	user3 //嵌入类型，内部类型
	level string
}

func (a *admin3) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func testNestStruct() {

	ad := admin3{
		user3: user3{
			name:  "aaa",
			email: "xxcvxc",
		},
		level: "super",
	}
	sendNotification3(&ad)

	//访问内部类的方法
	ad.user3.notify()
	//访问外部类或内部类提升来的方法
	ad.notify()

}

func sendNotification3(n notifier3) {
	n.notify()
}
