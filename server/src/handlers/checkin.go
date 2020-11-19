package handlers

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"encoding/json"

	"fmt"
	"io/ioutil"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type CheckInfo struct {
	time  string
	class string
	id    string
	tp    string
	ctime string
	pic   string
}

func (a CheckInfo) toString() string {
	var res = a.time + "," + a.class + "," + a.id + "," + a.tp + "," + a.ctime

	return res
}

func (a CheckInfo) toString_sql() string {
	var res = "('" + a.time + "','" + a.class + "','" + a.id + "','" + a.tp + "','" + a.ctime + "','')"
	return res
}

func resolveCheckInfo(q [][]sql.RawBytes) string {
	var r []string
	for _, a := range q {
		t := *(*string)(unsafe.Pointer(&a[0]))
		c := *(*string)(unsafe.Pointer(&a[1]))
		tp := *(*string)(unsafe.Pointer(&a[2]))
		r = append(r, t+","+c+","+tp)
	}
	return strings.Join(r, ";")
}

func resolveClassNameTime(q [][]sql.RawBytes) string {
	var r []string
	for _, a := range q {
		name := *(*string)(unsafe.Pointer(&a[0]))
		te := *(*string)(unsafe.Pointer(&a[1]))
		r = append(r, name+","+te)
	}
	return strings.Join(r, ";")
}

func resolveIdList(q [][]sql.RawBytes) []string {
	var r []string
	for _, a := range q {
		id := *(*string)(unsafe.Pointer(&a[0]))
		r = append(r, id)
	}
	return r
}

func resolveTimeCount(q [][]sql.RawBytes) string {
	var r []string
	for _, a := range q {
		ddl := *(*string)(unsafe.Pointer(&a[0]))
		stime := *(*string)(unsafe.Pointer(&a[1]))
		count := *(*string)(unsafe.Pointer(&a[2]))
		r = append(r, ddl+","+stime+","+count)
	}
	return strings.Join(r, ";")
}

func resolveNameIdTypeTime(q [][]sql.RawBytes, tp string) string {
	var res string
	var ending string
	if tp == "download" {
		ending = "\n"
	} else {
		ending = ";"
	}
	for _, a := range q {
		name := *(*string)(unsafe.Pointer(&a[0]))
		id := *(*string)(unsafe.Pointer(&a[1]))
		tp := changeType(*(*string)(unsafe.Pointer(&a[2])))
		ctime := changeTimeType(*(*string)(unsafe.Pointer(&a[3])))
		res += name + "," + id + "," + tp + "," + ctime + ending
	}
	return res
}

func changeType(s string) string {
	if s == "0" {
		return "未签到"
	} else if s == "1" {
		return "签到码签到"
	} else if s == "2" {
		return "人脸识别签到"
	} else {
		return s
	}
}

func changeTimeType(s string) string {
	if s == "" {
		return s
	}
	return s[:4] + "/" + s[4:6] + "/" + s[6:8] + " " + s[8:10] + ":" + s[10:]
}

func resolveTimeType(q [][]sql.RawBytes) string {
	var res []string
	for _, a := range q {
		tp := *(*string)(unsafe.Pointer(&a[0]))
		ctime := *(*string)(unsafe.Pointer(&a[1]))
		res = append(res, tp+","+ctime)
	}
	return strings.Join(res, ";")
}

func GetConfig(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var id, t string
	var d map[string]string

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		id = d["id"]
		t = d["t"]
	}

	var msg = "failed"
	var tp = ""
	var classList = ""
	var checkRes = ""

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
			classList = resolveClassName(qr3)
		} else if tp == "teacher" {
			qr4 := QueryFromDB(db, "select className, time from class_info where teacherId = '"+id+"';")
			classList = resolveClassNameTime(qr4)
		}

		qr5 := QueryFromDB(db, "select time,class,type from checkin where id = '"+id+"' and SUBSTRING(time, 1, 8) = '"+t[0:8]+"' and CAST(SUBSTRING(time, 9, 4) as SIGNED) >= "+t[8:]+";")
		checkRes = resolveCheckInfo(qr5)

		msg = "success," + "{\"classList\" : \"" + classList + "\", \"type\" : \"" + tp + "\", \"checkRes\" : \"" + checkRes + "\"}"
	}

	c.JSON(200, msg)
}

func GetCheckList(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var id, cname, tp string
	var d map[string]string

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		id = d["id"]
		cname = d["class"]
		tp = d["type"]
	}

	var msg = "failed"
	var result = ""

	open, db := OpenDB()
	if open {
		defer db.Close()

		if tp == "student" {
			qr1 := QueryFromDB(db, "select type,checkTime from checkin where class = '"+cname+"' and id = '"+id+"';")
			result = resolveTimeType(qr1)
		} else if tp == "teacher" {
			qr2 := QueryFromDB(db, "select a.time,b.checkTime,count(a.type) from checkin a, (SELECT time,checkTime FROM checkin where class = '"+cname+"' and id = '"+id+"') b where a.time = b.time and a.type ='0' group by checkTime,time;")
			qr3 := QueryFromDB(db, "select count(student) from class where className = '"+cname+"';")
			countList := resolveTimeCount(qr2)
			snum := *(*string)(unsafe.Pointer(&qr3[0][0]))
			result = countList + ";" + snum
		}

		msg = "success," + result
	}

	c.JSON(200, msg)
}

func Checkin(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var id string
	var cname string
	var ddl string
	var stime string
	var d map[string]string

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		id = d["id"]
		cname = d["cname"]
		ddl = d["ddl"]
		stime = d["stime"]
	}

	var msg = "failed"

	open, db := OpenDB()
	if open {
		defer db.Close()

		res := UpdateToDB(db, "update checkin set type = '1', checkTime = '"+stime+"' where id = '"+id+"' and class = '"+cname+"' and time = '"+ddl+"';")
		msg = "success," + string(res)
	}

	c.JSON(200, msg)
}

func OpenCheckin(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var id, cname, ddl, stime = "", "", "", ""
	var d map[string]string

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		id = d["id"]
		cname = d["cname"]
		ddl = d["ddl"]
		stime = d["stime"]
	}

	var msg = "failed"

	open, db := OpenDB()
	if open {
		defer db.Close()

		qr1 := QueryFromDB(db, "select student from class where className = '"+cname+"';")
		idList := resolveIdList(qr1)

		temp := CheckInfo{
			ddl,
			cname,
			id,
			"0",
			stime,
			"",
		}

		var values = temp.toString_sql() + ","
		temp.ctime = ""
		for i, a := range idList {
			temp.id = a
			idList[i] = temp.toString_sql()
		}
		values += strings.Join(idList, ",")

		qr2 := InsertToDB(db, "insert into checkin(time, class, id, type, checkTime, pic) values"+values+";")

		msg = "success," + string(qr2)
	}

	c.JSON(200, msg)
}

func CloseCheckin(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var cname, ddl = "", ""
	var d map[string]string

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		ddl = d["ddl"]
		cname = d["cname"]
	}

	var msg = "failed"

	open, db := OpenDB()
	if open {
		defer db.Close()

		qr := UpdateToDB(db, "delete from checkin where time = '"+ddl+"' and class = '"+cname+"';")
		os.RemoveAll("images/checkin/" + ddl + "_" + cname)
		msg = "success," + string(qr)
	}

	c.JSON(200, msg)
}

func count(db *sql.DB, cname string, ddl string) string {
	qr := QueryFromDB(db, "select count(id) from checkin where time = '"+ddl+"' and class = '"+cname+"' and type != '0';")
	return *(*string)(unsafe.Pointer(&qr[0][0]))
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetCount(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	var res = ""
	open, db := OpenDB()
	if open {
		defer db.Close()
	} else {
		res = "failed"
	}

	var cname, ddl = "", ""

	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}

		if res == "failed" {
			err = ws.WriteMessage(mt, []byte(res))
			break
		}

		var msg = string(message)
		fmt.Println(msg)

		if msg[0:4] == "open" {
			ml := strings.Split(msg, ",")
			cname, ddl = ml[1], ml[2]
		}

		if msg == "count" {
			res = count(db, cname, ddl)
			message = []byte(res)
			//写入ws数据
			err = ws.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}

	}
}

func Download(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var id, cname, te, tp string
	var d map[string]string

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		id = d["id"]
		cname = d["cname"]
		te = d["te"]
		tp = d["tp"]
	}

	var res string

	open, db := OpenDB()
	if open {
		defer db.Close()

		qr := QueryFromDB(db, "select b.name, a.id, a.type, a.checkTime from checkin as a, student_info as b where a.time = '"+te+"' and a.class = '"+cname+"' and a.id != '"+id+"' and b.id = a.id order by a.type desc;")
		res = resolveNameIdTypeTime(qr, tp)
	}

	c.JSON(200, res)
}

func GetFace(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var id, ddl, cname string
	var d map[string]string

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		id = d["id"]
		ddl = d["ddl"]
		cname = d["cname"]
	}

	var res string
	open, db := OpenDB()
	if open {
		defer db.Close()

		qr := QueryFromDB(db, "select pic from checkin where time = '"+ddl+"' and class = '"+cname+"' and id = '"+id+"';")
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
