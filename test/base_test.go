package test

import (
	"github.com/liyork/goinaction/mylib/fmt"
	"github.com/liyork/goinaction/prvate"
	"testing"
)

func TestBasic(t *testing.T) {

	//counter := fmt.alertCounter(10)//编译报错，未公开
	counter := fmt.AlertCounter1(10) //ok
	fmt.Println(counter)
	alertCounter := fmt.GetAlertCounter()
	fmt.Println(alertCounter)
	//必须满足：标识符不是一个值，使用短变量声明操作符(有能力捕获引用的类型，并创建一个未公开类型的变量)
	i := fmt.New(2)
	fmt.Println(i)
	u := prvate.Userx{
		Name: "aaa",
		//age:111,//不可见报错。
	}
	fmt.Println(u)
	//使用内部嵌入未公开但是由于提升，使用外部访问的属性
	a := prvate.Admin{
		Rights: 10,
	}
	a.Name = "biill"
	a.Email = "bill@ema.com"

}
