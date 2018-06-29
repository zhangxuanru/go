package main

import "fmt"

//空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。
//一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。

//空interface可以存储任意类型的数据
var a interface{}

func main()  {
	testInterface()

	switchInterface()
}

func testInterface()  {
	 var i int =  5
	  s := "hello world"
	  a = i
	  a = s
	  fmt.Println(a)
}

/**
  数组存储任意类型的值，并获取值的类型
 */
func switchInterface()  {
	  type Element interface {} //定义空接口
	  type Llist [] Element   //动态数组，可存储的类型是任意类型，因为Element是空interface
	  type Person struct {
	  	name string
	  	age int
	}
	list := make(Llist,3)   //创建3个空间
	list[0] = 1 //an int   赋值整型
	list[1] = "Hello" //a string  赋值字符串
	list[2] = Person{"Dennis", 70} // 赋值结构体

	for index,value := range list{
		switch  elen := value.(type){  //获取值的类型
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, elen)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}
