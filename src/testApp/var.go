package main

/*
定义变量 和常量
一般用var方式来定义全局变量
_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。
*/

//分组写成如下形式 start
	const(
		i = 100
		pi = 3.1415
		prefix = "Go_"
	)
	var(
		i int
		pi float32
		prefix string
	)
//分组写成如下形式 end

func main()  {
	var arr [10]int  // 声明了一个int类型的数组

	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}


	//定义一个名称为“variableName”，类型为"type"的变量
	var variableName type

	//定义三个类型都是“type”的变量
	var vname1, vname2, vname3 type

	//初始化“variableName”的变量为“value”值，类型是“type”
	var variableName type = value

	/*
	定义三个类型都是"type"的变量,并且分别初始化为相应的值
	vname1为v1，vname2为v2，vname3为v3
*/
	var vname1, vname2, vname3 type = v1, v2, v3

    //简化版
	vname1, vname2, vname3 := v1, v2, v3

	//_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。在这个例子中，我们将值35赋予b，并同时丢弃34：
	_, b := 34, 35


	// 定义常量
	const constantName = value
	//如果需要，也可以明确指定常量的类型：
	const Pi float32 = 3.1415926

	const Pi = 3.1415926
	const i = 10000
	const MaxThread = 10
	const prefix = "astaxie_"

}
