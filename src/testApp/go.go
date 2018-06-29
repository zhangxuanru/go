package main

import "fmt"

func main()  {
    var a  = []int{7, 2, 8, -9, 4, 0}
    var c = make(chan int)

	//ch:= make(chan bool, 4)，创建了可以存储4个元素的bool 型channel 前4个元素可以无阻塞的写入。当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间

	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)
    x,y := <-c, <-c
	fmt.Println(x, y, x + y)


    //range 不断的读取channel里面的数据
    ch := make(chan int,10)
	go rangeChannel(cap(ch),ch)
    for v := range ch{
		  fmt.Println(v)
	}
    fmt.Println(ch,cap(ch))
}


func sum(a []int, c chan int){
	sum := 0;
	for _, v:=range a{
		sum += v
	}
	c <- sum  // send sum to c
}

//range channel
func rangeChannel(n int,c chan int)  {
    x,y := 1,1
    for i:=0; i<n;i++{
    	c <-x
    	x,y = y,x+y
	}
	close(c)  //显式关闭channel
}


