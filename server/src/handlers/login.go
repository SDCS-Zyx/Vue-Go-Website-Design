package handlers

import (
	"database/sql"
	"encoding/json"

	// "fmt"
	"io/ioutil"
	"unsafe"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func check(u Account, res [][]sql.RawBytes) string {
	for _, a := range res {
		if *(*string)(unsafe.Pointer(&a[0])) == u.Password {
			return "success,id:" + *(*string)(unsafe.Pointer(&a[1]))
		}
	}

	return "wrong password"
}

func Login(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	var u Account

	if err := json.Unmarshal([]byte(string(data)), &u); err == nil {
		// fmt.Println(u.Username, u.Password)
	}

	open, db := OpenDB()
	var res string = "failed"
	if open {
		qr := QueryFromDB(db, "select password,id from account where username = '"+u.Username+"';")
		if len(qr) == 0 {
			res = "user not exist"
		} else {
			res = check(u, qr)
		}
		db.Close()
	}

	// fmt.Printf("body: %v", string(data))
	c.JSON(200, res)
}
