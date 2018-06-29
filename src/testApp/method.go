package main

import (
	"math"
	"fmt"
)

//定义一个长方形结构体
type Rectangle struct {
	 width,hight float64
}

//定义一个圆的结构体
type Circle struct {
	radius float64
}

//为长方形计算面积，面向对象形式
func (r Rectangle) area() float64  {
	  return  r.width * r.hight
}

//为圆计算面积，面向对象形式
func (c Circle) area() float64  {
   return  c.radius*c.radius*math.Pi
}

//////////////////////////////////////////////
type Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human //匿名字段
	school string
}
//为 Human 结构体加上 sayHi方法
func (h *Human) sayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//method重写，如果字段里已经写有方法， 但需要重写， 则直接写方法即可
func (s *Student) sayHi()  {
	fmt.Printf("Hi,重写sayHi方法 I am %s you can call me on %s\n", s.name, s.phone)
}

//应用匿名字段里定义的方法
func structFun()  {
	//实例 Student后， 也可以用 匿名字段里这定义的方法
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	mark.sayHi()
}



func main()  {
   r1 := Rectangle{10,20}
   r2 := Rectangle{9,4}
   c1 := Circle{10}
   c2 := Circle{20}
   fmt.Println("Area of r1 is: ", r1.area())
   fmt.Println("Area of r2 is: ", r2.area())
   fmt.Println("Area of c1 is: ", c1.area())
   fmt.Println("Area of c2 is: ", c2.area())

	structFun()
}



