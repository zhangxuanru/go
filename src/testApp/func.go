package main

import "fmt"

func main()  {
	m := max(10,20)
	xPLUSy, xTIMESy := SumAndProduct(10,20)
	x,y :=SumAndProduct2(10,20)
	//变参
   argParmenInt(6,7,90)

	xx := 10
	x1 := add1(&xx)
	fmt.Println(xx,x1)

   fmt.Println(m,xPLUSy,xTIMESy,x,y)
}

//多个返回值
func SumAndProduct(a,b int) (int,int) {
	 return a+b,a*b
}

//命名返回参数的变量,便于理解返回值的意思
func SumAndProduct2(a,b int)(add int,mi int)  {
	  add = a+b
	  mi  = a*b
	  return
}


//变参
 func argParmenInt(arg ...int) bool {
	for k,n := range arg{
		  fmt.Println(k,n)
	}
	return true
}

//传值与传指针
/*
传指针使得多个函数能操作同一个对象。
传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）
*/
func add1(a *int)  int  {
     *a = *a+1
     return *a
}

/**
比较两个数大小
 */
func max(a,b int)int {
	    if a > b{
	    	return  a
		}
  return b
}