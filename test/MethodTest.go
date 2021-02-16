package test

import (
	"fmt"
)

type user1 struct {
	name  string
	email string
}

//func与函数名之间的参数被称为[接收者]，将函数与接受者的类型绑在一起，如果一个函数有接收者，则称为方法。
//Go 语言里有两种类型的接收者： 值接收者和指针接收者
//如果使用值接收者声明方法，调用时会使用这个值的一个副本来执行。
//值接收者使用值的副本来调用方法，修改的是副本，而指针接受者使用实际值来调用方法，修改的相同值

//使用[值接收者]实现了一个方法,操作的是一个副本，绑定到了user1实例上
func (u user1) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n", u.name, u.email)
}

//使用[指针接收者]实现一个方法，这个方法会共享调用方法时接收者所指向的值
func (u *user1) changeEmail(email string) {
	u.email = email
}

func testMethod() {

	bill := user1{"bill","bill@mail.com"}
	//user类型的值可以调用-使用值接收者声明的方法
	//这个语法与调用一个包里的函数看起来很类似。但在这个例子里， bill 不是包名，而是变量名。
	bill.notify()

	//指针变量
	lisa := &user1{"lisa","lis@emai.com"}
	//指向user类型值的指针可以调用-使用值接收者声明的方法
	//编译器背后，(*lisa).notify()，指针被解引用为值
	lisa.notify()

	//user类型的值可以调用-使用指针接收者声明的方法
	//编译器背后，(&bill).changeEmail ("bill@newdomain.com")
	bill.changeEmail("bill@newdomain.com")
	//user类型值的指针可以调用-使用指针接收者声明的方法
	lisa.changeEmail("lisa@newdomain.com")

}
