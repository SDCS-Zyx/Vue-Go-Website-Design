package main

import (
	"router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.RouterInit(r)
	r.Run(":8081")
}
