package handlers

import (
	"bufio"
	"database/sql"
	"strings"

	"encoding/json"

	// "fmt"
	"io"
	"io/ioutil"
	"os"
	"unsafe"

	"github.com/gin-gonic/gin"
)

type ClassInfo struct {
	Name      string `json:"name"`
	School    string `json:"school"`
	Dep       string `json:"dep"`
	Time      string `json:"time"`
	Teacher   string `json:"teacher"`
	TeacherId string `json:"tid"`
}

func (a ClassInfo) toString() string {
	var res = a.Name + "," + a.School + "," + a.Dep + "," + a.Time + "," + a.Teacher + "," + a.TeacherId
	return res
}

func getName(q [][]sql.RawBytes) string {
	return *(*string)(unsafe.Pointer(&q[0][0]))
}

func resolveClassName(q [][]sql.RawBytes) string {
	var r []string
	for _, a := range q {
		now := "'" + *(*string)(unsafe.Pointer(&a[0])) + "'"
		r = append(r, now)
	}
	return strings.Join(r, ",")
}

func resolveClassInfo(q [][]sql.RawBytes) string {
	var r []string
	for _, a := range q {
		now := ClassInfo{
			*(*string)(unsafe.Pointer(&a[0])),
			*(*string)(unsafe.Pointer(&a[1])),
			*(*string)(unsafe.Pointer(&a[2])),
			*(*string)(unsafe.Pointer(&a[3])),
			*(*string)(unsafe.Pointer(&a[4])),
			*(*string)(unsafe.Pointer(&a[5])),
		}
		r = append(r, now.toString())
	}
	return strings.Join(r, ";")
}

func GetClassInfo(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var d map[string]string
	var id string

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		id = d["id"]
	}

	var msg = "failed"
	var tp = ""
	var result = ""

	open, db := OpenDB()
	if open {
		defer db.Close()

		qr1 := QueryFromDB(db, "select name from student_info where id = '"+id+"';")
		if len(qr1) == 0 {
			qr2 := QueryFromDB(db, "select name from teacher_info where id = '"+id+"';")
			if len(qr2) != 0 {
				tp = "teacher"
			}
		} else {
			tp = "student"
		}

		if tp == "student" {
			qr3 := QueryFromDB(db, "select className from class where student = '"+id+"';")
			name_list := resolveClassName(qr3)
			qr4 := QueryFromDB(db, "select * from class_info where className in ("+name_list+");")
			result = resolveClassInfo(qr4)
		} else if tp == "teacher" {
			qr5 := QueryFromDB(db, "select * from class_info where teacherId = '"+id+"';")
			result = resolveClassInfo(qr5)
		}

		msg = "success," + result + ";" + tp
	}

	c.JSON(200, msg)
}

func dealFile(filename string) string {
	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	sl := strings.Split(filename, "_")
	cname := sl[0]
	tid := sl[1]
	ctime := strings.Split(sl[2], ".")[0]

	open, db := OpenDB()
	if !open {
		return "failed"
	}
	defer db.Close()

	qr1 := QueryFromDB(db, "select * from teacher_info where id = '"+tid+"';")
	if len(qr1) == 0 {
		return "wrong id"
	}
	var tinfo TeacherInfo
	resolveTeacherInfo(&tinfo, qr1)

	tx, _ := db.Begin()

	tx.Exec("insert into class_info(className, school, department, time, teacher, teacherId) values(?,?,?,?,?,?);", cname, tinfo.School, tinfo.Dep, ctime, tinfo.Name, tid)

	br := bufio.NewReader(file)
	for {
		line, _, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1 //读取出错
			}
			//读取到文件最后
			break
		}

		//isPrefix : true 表示一行数据字节太长

		sid := string(line)
		tx.Exec("insert into class(className, student) values(?,?);", cname, sid)
	}
	tx.Commit()

	return "success"
}

func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	checkErr(err)

	filename := header.Filename
	out, err := os.Create(filename)
	checkErr(err)
	defer out.Close()

	_, err = io.Copy(out, file)
	checkErr(err)

	res := dealFile(filename)

	c.String(200, res)
}
