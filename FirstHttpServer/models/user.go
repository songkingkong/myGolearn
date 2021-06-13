package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserName string `orm:"size(100)"`
	PassWd   string `orm:"size(100)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)
	orm.RegisterModel(new(User))
	orm.SetMaxOpenConns("default", 30)
	orm.SetMaxIdleConns("default", 10)
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}
func (u *User) AddUser(user string, passwd string) {
	o := orm.NewOrm()
	u.UserName = user
	u.PassWd = passwd
	_, err := o.Insert(&u)
	if err != nil {
		fmt.Println(err)
	}
}
func (u *User) SelectUser(user string) (string, error) {
	o := orm.NewOrm()
	u.UserName = user
	err := o.Read(&u)
	return u.PassWd, err
}
