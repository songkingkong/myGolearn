package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"myGolearn/testserver/models"
	"net/http"
	"strings"
)

type JsonResult struct {
	Code int `json:"code"`
	Msg  models.User
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
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

		u := models.User{}
		u.UserName = strings.Join(r.Form["username"], "")
		u.PassWd = strings.Join(r.Form["password"], "")

		msg := JsonResult{Msg: u}

		userPasswd, err := u.SelectUser(u.UserName)
		if err != nil {
			fmt.Println(err)
		}
		if u.PassWd == userPasswd {
			msg.Code = 200
		} else {
			msg.Code = 400
		}

		result, _ := json.Marshal(msg)
		w.Header().Set("content-type", "text/json")
		w.WriteHeader(200)
		w.Write(result)
	}
}
func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
