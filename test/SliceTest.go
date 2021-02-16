package test

import "fmt"

func testCreate() {
	//创建
	slice1 := make([]string, 5) //长度、容量都是5
	slice2 := make([]int, 3, 5) //长度3、容量5，长度不允许大于容量

	slice3 := []string{"red", "blue", "3", "4", "5"}
	slice4 := []int{11, 2, 3}

	slice5 := []string{"99", ""} //初始化第99元素

	fmt.Print(slice1, slice2, slice3, slice4, slice5)
}

func testDiff() {
	//如果在[]运算符里指定了一个值，那么创建的就是数组而不是切片。只有不指定值的时候，才会创建切片
	array := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}

	fmt.Print(array, slice)
}

func testNil() {
	//nil切片，描述不存在的切片，例如异常
	var slice1 []int

	//空切片，底层数组包含0个元素，没有分配任何存储空间。表示空集合，如数据库查询返回0个查询结果
	slice2 := make([]int, 0)
	slice3 := []int{}

	fmt.Print(slice1, slice2, slice3)
}

func testUpdateSlice() {
	slice := []int{1, 2, 3, 4, 5}
	//赋值
	slice[1] = 25
}

func testLen() {
	//切片之所以被称为切片，是因为创建一个新的切片就是把底层数组切出一部分
	//两个切片共享同一个底层数组。如果一个切片修改了该底层数组的共享部分，另一个切片也能感知到
	//整形切片，长度和容量5
	slice := []int{1, 2, 3, 4, 5}
	//新切片，长度为2，容量为4，起始位置为1
	newSlice := slice[1:3]

	//计算长度和容量：
	//底层数组容量是k的切片slice[i:j]则长度=j-i，容量=k-i

	//切片只能访问到其长度内的元素。试图访问超出其长度的元素将会导致语言运行时异常
	//newSlice[3] = 33 //运行错误

	fmt.Print(newSlice)
}

func testAppend() {
	//相对于数组而言，使用切片的一个好处是，可以按需增加切片的容量，用append函数
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3]
	//使用原有容量来分配一个新元素，新元素赋值为6，新切片长度为3，容量为4
	newSlice = append(newSlice, 6)

	//函数 append 会智能地处理底层数组的容量增长。在切片的容量小于 1000 个元素时，总是
	// 会成倍地增加容量。一旦元素个数超过 1000，容量的增长因子会设为 1.25，也就是会每次增加 25%的容量
	newSlice2 := append(slice, 6)
	fmt.Print(newSlice2)
}

func testLimit() {
	//限制容量
	source := []string{"1", "2", "3", "4", "5"}
	//将第三个元素切片，并限制容量，长度1，容量为2
	slice := source[2:3:4]
	fmt.Print(slice)

	//slice[i:j:k]则长度=j-i,容量=k-i
	//内置函数 append 会首先使用可用容量。一旦没有可用容量，会分配一个新的底层数组。
	slice2 := source[2:3:3]      //长度1，容量1
	slice2 = append(slice2, "4") //没有容量，则新构造数组拷贝数据，添加4

	//合并切片
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...)) //s2追加到s1后面
}

func testSlice() {

	//range返回索引和对应值的副本(不是原值引用)，就不用&value这样使用指针地址方式获取值
	//关键字 range 总是会从切片头部开始迭代
	slice1 := []int{1, 2, 3, 4}
	for index, value := range slice1 {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	//for
	// 对于切片，函数 len 返回切片的长度，函数 cap 返回切片的容量
	slice := []int{1, 2, 3, 4}
	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}

	//多维切片
	//外层切片包含两个元素，每个元素都是一个切片，
	slice2 := [][]int{{10}, {100, 200}}

	//为第一个切片追加值为20的元素
	slice2[0] = append(slice2[0], 20)

	//传递切片
	//在 64 位架构的机器上，一个切片需要 24 字节的内存：指针字段需要 8 字节，长度和容量字段分别需要 8 字节
	//复制时只会复制切片本身，不会涉及底层数组
	//在函数间传递 24 字节的数据会非常快速、简单。这也是切片效率高的地方
	slice3 := make([]int, 1000)
	fmt.Print(slice3)
	slice = transmitSlice(slice)
}

func transmitSlice(slice []int) []int {
	return slice
}
