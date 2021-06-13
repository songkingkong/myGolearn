package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strings"
)

/*import (
	_ "FirstHttpServer/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}*/

type UserInfo struct {
	UserName string
	PassWord string
}
type JsonResult struct {
	Code int `json:"code"`
	Msg  UserInfo
}
type User struct {
	UserName string `orm:"PK"`
	PassWd   string `orm:"size(100)"`
}

func (u *User) SelectUser(user string) (string, error) {
	o := orm.NewOrm()
	u.UserName = user
	err := o.Read(&u)
	return u.PassWd, err
}
func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)
	orm.RegisterModel(new(User))
	orm.SetMaxOpenConns("default", 30)
	orm.SetMaxIdleConns("default", 10)
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			fmt.Println(err)
		}
		log.Println(t.Execute(w, nil))
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			fmt.Println(err)
		}
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm:", err)
		}
		username := strings.Join(r.Form["username"], "")
		password := strings.Join(r.Form["password"], "")
		o := orm.NewOrm()
		u := User{UserName: username}
		err = o.Read(&u)
		userinfo := UserInfo{UserName: username, PassWord: password}
		msg := JsonResult{Msg: userinfo}
		if err != nil {
			fmt.Println(err)
		}
		if password == u.PassWd {
			msg.Code = 200
		} else {
			msg.Code = 400
		}
		w.Header().Set("content-type", "text/json")
		result, _ := json.Marshal(msg)
		w.WriteHeader(200)
		w.Write(result)
	}
}
func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
