package handlers

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"unsafe"

	"github.com/gin-gonic/gin"
)

type TeacherInfo struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	Type   string `json:"type"`
	School string `json:"school"`
	Dep    string `json:"dep"`
	Sex    string `json:"sex"`
}

func (a Info) toString() string {
	var res = a.Name + "," + a.Id + "," + a.Type + "," + a.School + "," + a.Dep + "," + a.Subject + "," + a.Date1 + "," + a.Date2 + "," + a.Sex
	return res
}

func (a TeacherInfo) toString() string {
	var res = a.Name + "," + a.Id + "," + a.Type + "," + a.School + "," + a.Dep + "," + a.Sex
	return res
}

func resolveInfo(s *Info, qr [][]sql.RawBytes) {
	var i = qr[0]
	s.Name = *(*string)(unsafe.Pointer(&i[0]))
	s.Id = *(*string)(unsafe.Pointer(&i[1]))
	s.Type = "student"
	s.School = *(*string)(unsafe.Pointer(&i[2]))
	s.Dep = *(*string)(unsafe.Pointer(&i[3]))
	s.Subject = *(*string)(unsafe.Pointer(&i[4]))
	s.Date1 = *(*string)(unsafe.Pointer(&i[5]))
	s.Date2 = *(*string)(unsafe.Pointer(&i[6]))
	s.Sex = *(*string)(unsafe.Pointer(&i[7]))
}

func resolveTeacherInfo(s *TeacherInfo, qr [][]sql.RawBytes) {
	var i = qr[0]
	s.Name = *(*string)(unsafe.Pointer(&i[0]))
	s.Id = *(*string)(unsafe.Pointer(&i[1]))
	s.Type = "teacher"
	s.School = *(*string)(unsafe.Pointer(&i[2]))
	s.Dep = *(*string)(unsafe.Pointer(&i[3]))
	s.Sex = *(*string)(unsafe.Pointer(&i[4]))
}

func GetInfo(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var d map[string]string
	var id = "-1"

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		id = d["id"]
	}

	var msg = "failed"

	open, db := OpenDB()
	if open {
		qr := QueryFromDB(db, "select * from student_info where id = '"+id+"';")
		if len(qr) != 0 {
			var s Info
			resolveInfo(&s, qr)
			fmt.Println(s)
			msg = "success," + s.toString()
		} else {
			qr2 := QueryFromDB(db, "select * from teacher_info where id = '"+id+"';")
			if len(qr2) != 0 {
				var t TeacherInfo
				resolveTeacherInfo(&t, qr2)
				fmt.Println(t)
				msg = "success," + t.toString()
			}
		}
		db.Close()
	}

	c.JSON(200, msg)
}

func GetPic(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var id string
	var d map[string]string

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		id = d["id"]
	}

	var res string
	open, db := OpenDB()
	if open {
		defer db.Close()

		qr := QueryFromDB(db, "select pic from user_pic where id = '"+id+"';")
		if len(qr) != 0 {
			filePath := *(*string)(unsafe.Pointer(&qr[0][0]))
			img, err := ioutil.ReadFile(filePath)
			checkErr(err)
			imgBase64 := base64.StdEncoding.EncodeToString(img)

			res = imgBase64
		}
	}

	c.JSON(200, res)
}

func ChangeInfo(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	fmt.Println(data)

	var s map[string]Info
	var d Info

	if err := json.Unmarshal([]byte(string(data)), &s); err == nil {
		d = s["info"]
	}

	var res = "failed"
	open, db := OpenDB()
	defer db.Close()
	if open {
		if d.Type == "student" {
			r := UpdateToDB(db, "update student_info set name='"+d.Name+"',school='"+d.School+"',department='"+d.Dep+"',subject='"+d.Subject+"',date1='"+d.Date1+"',date2='"+d.Date2+"',sex='"+d.Sex+"' where id='"+d.Id+"';")
			res = "success," + string(r)
		} else if d.Type == "teacher" {
			r := UpdateToDB(db, "update teacher_info set name='"+d.Name+"',school='"+d.School+"',department='"+d.Dep+"',sex='"+d.Sex+"' where id='"+d.Id+"';")
			res = "success," + string(r)
		}
	}

	c.JSON(200, res)
}
