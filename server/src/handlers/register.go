package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "strconv"

	"github.com/gin-gonic/gin"
)

type Info struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	School  string `json:"school"`
	Dep     string `json:"dep"`
	Subject string `json:"subject"`
	Date1   string `json:"date1"`
	Date2   string `json:"date2"`
	Sex     string `json:"sex"`
}

type postData struct {
	Form    Info    `json:"form"`
	Account Account `json:"account"`
}

func Register(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var con postData
	var u Info
	var a Account

	if err := json.Unmarshal([]byte(string(data)), &con); err == nil {
		u = con.Form
		a = con.Account
		fmt.Println(u, a)
	}

	open, db := OpenDB()
	tx, _ := db.Begin()
	var res string = "failed"
	if open {
		//多个数据库操作时，若有其中一个操作出错，所有操作需要roll back
		tx.Exec("insert into account(username, password, id) values('" + a.Username + "', '" + a.Password + "', '" + u.Id + "');")

		if u.Type == "student" {
			tx.Exec("insert into student_info(name,id,school,department,subject,date1,date2,sex) values('" + u.Name + "', '" + u.Id + "', '" + u.School + "', '" + u.Dep + "', '" + u.Subject + "', '" + u.Date1 + "', '" + u.Date2 + "', '" + u.Sex + "');")
		} else if u.Type == "teacher" {
			tx.Exec("insert into teacher_info(name,id,school,department,sex) values('" + u.Name + "', '" + u.Id + "', '" + u.School + "', '" + u.Dep + "', '" + u.Sex + "');")
		}

		tx.Commit()

		res = "success"
		db.Close()
	}
	c.JSON(200, res)
}

func CheckAccount(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	var d map[string]string

	var tryName = ""

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		tryName = d["username"]
		fmt.Println(tryName)
	}

	open, db := OpenDB()
	var res string = "failed, username exit"
	if open {
		qr := QueryFromDB(db, "select * from account where username = '"+tryName+"'")
		if len(qr) == 0 {
			res = "not use"
		}
		db.Close()
	}
	c.JSON(200, res)
}

func CheckId(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	var d map[string]string

	var id = ""
	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		id = d["id"]
	}

	open, db := OpenDB()
	var res string = "failed, id exit"
	if open {
		qr := QueryFromDB(db, "select id from student_info where id = '"+id+"'")
		qr1 := QueryFromDB(db, "select id from teacher_info where id = '"+id+"'")
		if len(qr) == 0 && len(qr1) == 0 {
			res = "not use"
		}
		db.Close()
	}
	c.JSON(200, res)
}
