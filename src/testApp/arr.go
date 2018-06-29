package main

import "fmt"

/*
slice的index只能是｀int｀类型

对于slice有几个有用的内置函数：

len 获取slice的长度
cap 获取slice的最大容量
append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数

*/
func main()  {
	sliceArr()
	runSlice()
}

//slice 动态数组
func sliceArr()  {
	   arr := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	   var a,b []byte
	   a = arr[2:5]
	   b = arr[3:5]
	   fmt.Println(a,b)
	   fmt.Println(arr[:]) //获取数组所有值
}

func runSlice()  {
    var arr []byte
	arr = append(arr,90)
	arr = append(arr,80)
	arr = append(arr,70)
	arr = append(arr,60)
	arr = append(arr,50)
	arr = append(arr,40)
	arr = append(arr,30)
	arr = append(arr,20)
	arr = append(arr,10)
	arr = append(arr,0)
	fmt.Println(arr,cap(arr),len(arr))
}

