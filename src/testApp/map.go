package main

import "fmt"

/*
MAP
map的index多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型

map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
map的长度是不固定的，也就是和slice一样，也是一种引用类型
内置的len函数同样适用于map，返回map拥有的key的数量
map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制

map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变
*/

func main()  {
	runMap()
	initMap()
	buiMap()
}

func runMap()  {
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	// var numbers map[string]int

	// 声明一个key是字符串，值为int的字典,
	 var  numbers = make(map[string]int)
	 numbers["one"] = 1  //赋值
	 numbers["ten"] = 10 //赋值
	 numbers["three"] = 3
	 fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
}

func initMap()  {
	// 初始化一个字典,直接初始化不用make
	rating := map[string]float32{"C":5,"Go":4.5,"Python":4.5,"C++":2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating,ok := rating["C#"]

	if ok{
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	}else{
		fmt.Println("We have no rating associated with C# in the map")
	}
	// 删除key为C的元素
	delete(rating,"C")
}

//map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变
func buiMap()  {
	 m := make(map[string]string)
	 m["Hello"] = "world"
	 fmt.Println(m)
	 m1 := m
  	 m1["Hello"] = "Salut"
	 fmt.Println(m1,m)
}





