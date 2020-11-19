package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
)

func FaceCheckin(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var id, img string
	var d map[string]string

	if err := json.Unmarshal([]byte(string(data)), &d); err == nil {
		id = d["id"]
		img = d["img"]
	}

	var t = GetTime()
	var res = "failed"

	open, db := OpenDB()
	if open {
		defer db.Close()

		qr1 := QueryFromDB(db, "select time,class,type from checkin where id = '"+id+"' and SUBSTRING(time, 1, 8) = '"+t[0:8]+"' and CAST(SUBSTRING(time, 9, 4) as SIGNED) >= "+t[8:]+";")
		if len(qr1) != 0 {
			ddl := *(*string)(unsafe.Pointer(&qr1[0][0]))
			cname := *(*string)(unsafe.Pointer(&qr1[0][1]))
			tp := *(*string)(unsafe.Pointer(&qr1[0][2]))

			if tp != "2" {
				imgPath := saveImage(ddl, cname, id, img)
				qr2 := UpdateToDB(db, "update checkin set type = '2',checkTime = '"+t+"',pic = '"+imgPath+"' where id = '"+id+"' and time = '"+ddl+"' and class ='"+cname+"';")
				res = "success," + string(qr2)
			}
		}

	}

	c.JSON(200, res)
}

func GetTime() string {
	var t = time.Now()
	var timeLayoutStr = "2006-01-02 15:04:05"
	ts := t.Format(timeLayoutStr)[:16]
	ts = strings.ReplaceAll(ts, "-", "")
	ts = strings.ReplaceAll(ts, ":", "")
	ts = strings.Replace(ts, " ", "", 1)
	return ts[:12]
}

func saveImage(ddl string, cname string, id string, img string) string {
	var enc = base64.StdEncoding
	path := "images/checkin/" + ddl + "_" + cname

	if err := os.Mkdir(path, os.ModePerm); err != nil {
		fmt.Println(err)
	}

	if img[11] == 'j' {
		img = img[23:]
		path += "/" + id + ".jpg"
	} else if img[11] == 'p' {
		img = img[22:]
		path += "/" + id + ".png"
	} else if img[11] == 'g' {
		img = img[22:]
		path += "/" + id + ".gif"
	}

	data, err1 := enc.DecodeString(img)
	checkErr(err1)

	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(data)

	return path
}
