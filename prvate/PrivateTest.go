package prvate

type Userx struct {
	Name string //公开
	age  int    //未公开
}

//2内部提升
type usery struct { //非公开
	Name  string //公开
	Email string //公开
}

type Admin struct {
	usery      //嵌入未公开
	Rights int //公开
}


