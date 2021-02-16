package main

import (
	"fmt"

	myfmt "github.com/liyork/goinaction/mylib/fmt" //重命名导入

	_ "database/sql" //使用空白标识符导入包，让初始化，避免编译错误
)

//$GOPATH/src/hello/中执行go build/go clean
//执行./hello
func main() {

	fmt.Println("Hello World!")

	myfmt.Println("Hello World!")
}
