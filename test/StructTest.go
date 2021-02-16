package test

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

//基于已有类型，自定义类型,从内置类型创建出很多更加明确的类型，并赋予更高级的功能
//见time.go中Duration，Duration 类型依然是一个独立的类型，两种不同类型的 值即便互相兼容，也不能互相赋值
type Duration int64

func testStruct() {
	//声明user类型的变量
	//var创建类类型为user且别名为bill的变量
	var bill user
	fmt.Print(bill)

	//:=，短变量声明操作符，一个短变量声明操作符在一次操作中完成两件事情：声明一个变量，并初始化
	//短变量 声明操作符会使用右侧给出的类型信息作为声明变量的类型
	lisa1 := user{
		name:       "lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}
	fmt.Print(lisa1)

	lisa2 := user{"lisa", "lis@emai.cmo", 123, true}
	fmt.Print(lisa2)

	//使用结构字面量来创建字段的值
	fred3 := admin{
		person: user{
			name:       "lisa",
			email:      "ls@ema.cmo",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}
	fmt.Print(fred3)
}
