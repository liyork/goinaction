package test

import "fmt"

func testDeclare() {
	//声明一个数组，并设置为零值
	var array [5]int
	fmt.Print(array)

	//使用数组字面量声明数组
	//array := [5]int{1, 2, 3, 4, 5}

	//自动计算长度
	//array := [...]int{1, 2, 3, 4, 5}

	//array := [5]int{1: 1, 2: 2}
}

func testUpdate() {
	//声明一个数组，并设置为零值
	var array [5]int

	//修改
	array[2] = 5
}

func testPointArray() {
	//包含指针的数组
	array := [5]*int{0: new(int), 1: new(int)}
	*array[0] = 10
	*array[1] = 20
}

func testArray() {

	//数组变量的类型包括数组长度和每个元素的类型。只有这两部分都相同的数组，才是类型相同的数组，才能互相赋值
	var array1 [5]string
	array2 := [5]string{"Red", "Blue", "Green", "Yello", "Pink"}
	array1 = array2
	fmt.Print(array1)

	var array3 [3]*string
	var array4 = [3]*string{new(string), new(string), new(string)}
	fmt.Print(array3)

	*array4[0] = "red"
	*array4[1] = "blue"
	*array4[2] = "yellow"
	//把一个指针数组赋值给另一个，只会复制指针的值，而不会复制指针所指向的值
	array3 = array4

}

func testMuliArray() {
	//二维数组
	var array1 [4][2]int
	array2 := [4][2]int{{1, 2}, {3, 4}, {5, 7}, {9, 0}}
	array3 := [4][2]int{1: {1, 2}, 3: {3, 4}}
	fmt.Print(array1)
	fmt.Print(array2)
	fmt.Print(array3)

	var array [2][2] int
	array[0][0] = 10
	array[0][1] = 20
	array[1][0] = 30
	array[1][1] = 40
}

//拷贝数组内的指针
func operatr(array *[1000]int) {

}
