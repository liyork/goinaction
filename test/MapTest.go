package test

import "fmt"

func testMap() {

	//创建一个映射，键类型string，值类型int
	dict1 := make(map[string]int)
	fmt.Print(dict1)
	//key:string,value:string
	dict2 := map[string]string{"1": "red", "2": "xdsf", "3": "1", "4": "sss"}
	fmt.Print(dict2)

	//映射的键可以是任何值。类型可以是内置的类型，也可以是结构类型，只要这个值
	//可以使用==运算符做比较。切片、函数以及包含切片的结构类型这些类型由于具有引用语义，
	//不能作为映射的键，使用这些类型会造成编译错误
	//dict := map[[]string]int{} //错误

	//值可以任意
	dict3 := map[int][]string{}
	fmt.Print(dict3)

	colors := map[string]string{}
	//赋值
	colors["red"] = "#da1111"

	//nil映射，不能存储键值，
	var color4 map[string]string
	fmt.Print(color4)

	//判断存在1
	value, exists := colors["blue"]
	if exists {
		fmt.Print(value)
	}

	//判断存在2，通过键来索引映射时，即便这个键不存在也总会返回一个值。在这种情况下，返回的是该值对应的类型的零值
	value2 := colors["blue"]
	if value2 != "" {
		fmt.Printf(value2)
	}

	//迭代
	colors3 := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	for key, value := range colors3 {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	delete(colors, "a")

	//在函数间传递映射并不会制造出该映射的一个副本。实际上，当传递映射给一个函数，并对这个映射做了修改时,
	//这个特性和切片类似，保证可以用很小的成本来复制映射
	removeColor(colors, "b")

	var lx = len(colors)
	fmt.Print(lx)
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
