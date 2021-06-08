package variables

import (
	"fmt"
	"math"
)

func Run() {
	// 定义一个变量，语法为var name type,这里int变量age没有初始值，默认值为0,string类型默认值为“”
	var age int
	//变量如果定义的时候有初始值，go会根据初始值推导出其类型，所以可以省略type
	var var1 = 10

	//同时定义多个变量,类型必须一致，不能同时定义不同类型的变量
	var width1, height1 int
	//同样，如果有初始值，则可以省略type,但是多个变量必须都有初始值
	var width2, height2 = 100, 50
	//同时定义多个不同类型的变量
	var (
		length = 1
		name   = "initialvalue1"
	)

	//简写变量的定义，语法为name := initialvalue
	//简写定义多个变量时，可以定义多个不同类型的变量，但是必须至少有一个新的变量并且都有初始值，否则报错
	count := 10
	name1, age := "newValue", math.MinInt64

	//如果想给多个已经定义的变量重新赋值，则可以使用如下语法：
	name1, age = "newValue1", 100

	fmt.Println("没有初始值的int变量默认值为", age)
	fmt.Println("有初始值的int变量值为", var1)
	fmt.Println("同时定义同一个类型的多个变量", width1, height1, width2, height2)
	fmt.Println("同时定义不同类型的多个变量", length, name)
	fmt.Println("使用简写同时定义多个变量", count, name1, age)
}
