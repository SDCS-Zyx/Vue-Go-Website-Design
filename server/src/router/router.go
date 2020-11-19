package router

import (
	// "net/http"

	"handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterInit(router *gin.Engine) {
	router.Use(cors.Default())

	router.POST("/Login", handlers.Login)

	router.POST("/Register/Register", handlers.Register)
	router.POST("/Register/checkAccount", handlers.CheckAccount)
	router.POST("/Register/checkId", handlers.CheckId)

	router.POST("/Home/profile/getInfo", handlers.GetInfo)
	router.POST("/Home/profile/getPic", handlers.GetPic)
	router.POST("/Home/profile/changeInfo", handlers.ChangeInfo)

	router.POST("/Home/classInfo/getClassInfo", handlers.GetClassInfo)

	router.POST("/Home/checkin/getConfig", handlers.GetConfig)
	router.POST("/Home/checkin/getCheckList", handlers.GetCheckList)
	router.POST("/Home/checkin/checkin", handlers.Checkin)
	router.POST("/Home/checkin/openCheckin", handlers.OpenCheckin)
	router.POST("/Home/checkin/closeCheckin", handlers.CloseCheckin)
	router.GET("/Home/checkin/getCount", handlers.GetCount)
	router.POST("/Home/checkin/getFace", handlers.GetFace)

	router.POST("/upload", handlers.Upload)
	router.POST("/download", handlers.Download)

	router.POST("/face_checkin", handlers.FaceCheckin)
}
