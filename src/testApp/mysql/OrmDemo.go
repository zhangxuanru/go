package main

/*
ORM:
https://beego.me/docs/mvc/model/orm.md

https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.5.md
*/
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	Id          int
	Name        string
	Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`    //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}



func main()  {
	insertData()
}

func init()  {
	orm.RegisterDataBase("default","mysql","root:abc123123@tcp(127.0.0.1:3306)/test?charset=utf8")
    orm.RegisterModel(new(User),new(Post),new(Profile),new(Tag))
	orm.RunSyncdb("default",false,true)
	orm.SetMaxIdleConns("default",30)
	orm.SetMaxOpenConns("default",30)
	orm.Debug =  true
}

func insertData()  {
	o := orm.NewOrm()
	o.Using("default")
	profile := new(Profile)
	profile.Age = 30
	user := new(User)
	user.Profile = profile
	user.Name="zhangxuanru"
	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}




