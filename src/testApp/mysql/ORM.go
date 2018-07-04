package main

import (
	 "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

// Model Struct
type User struct {
	Id int
	Name string `orm:"size(100)"`
}

//ORM 增删改查
/*
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.5.md

https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.5.md
*/
func  main()  {
	 insert()
	 del()
	 update()
	 getALL()
}

func init()  {
   orm.RegisterDataBase("default","mysql","root:abc123123@tcp(23.106.155.177:3306)/test?charset=utf8")
   orm.RegisterModel(new(User))
   orm.RunSyncdb("default",false,true)
   //根据数据库的别名，设置数据库的最大空闲连接
   orm.SetMaxIdleConns("default",30)
   //根据数据库的别名，设置数据库的最大数据库连接
   orm.SetMaxOpenConns("default",30)
   //目前beego orm支持打印调试，你可以 打开调试
   orm.Debug = true
}

func insert()  {
	o := orm.NewOrm()
	user := User{Name:"zxr"}
	id,err := o.Insert(&user)
	fmt.Printf("id = %d,user=%v,Err = %v\n",id,user,err)
}

func update()  {
	 o := orm.NewOrm()
	 user := User{Id:4,Name:"hello world"}
	 num,err := o.Update(&user)
	 fmt.Printf("Num:%d,user:%v,err:%v\n",num,user,err)
}

func del()  {
	o := orm.NewOrm()
	user := User{Id:2}
	num,err := o.Delete(&user)
	fmt.Printf("num:%d,user:%v,err=%v\n",num,user,err)
}

func getALL()  {
	 o := orm.NewOrm()
	 user := User{Id:3}
	 err := o.Read(&user)
	 fmt.Printf("user:%v,err:%v",user,err)
}


