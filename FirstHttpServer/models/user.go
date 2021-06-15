package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserName string `json:"user_name" orm:"PK" orm:"size(100)"`
	PassWd   string `json:"pass_wd" orm:"size(100)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)
	orm.RegisterModel(new(User))
	orm.SetMaxOpenConns("default", 30)
	orm.SetMaxIdleConns("default", 10)
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}
func (u *User) AddUser(username string, passwd string) error {
	o := orm.NewOrm()
	user := User{
		UserName: username,
		PassWd:   passwd,
	}
	_, err := o.Insert(&user)
	if err != nil {
		return err
	}
	return err
}
func (u *User) SelectUser(username string) (string, error) {
	o := orm.NewOrm()
	user := User{UserName: username}
	err := o.Read(&user)
	return user.PassWd, err
}
