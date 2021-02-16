package main

//引用
import (
	"log"
	"os"

	_ "github.com/liyork/goinaction/sample/matchers" //本文件未使用，但是进行初始化
	"github.com/liyork/goinaction/sample/search"
)

//在main之前执行
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president") //调用
}
