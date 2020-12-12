package main

import (
	"github.com/gin-gonic/gin"
)

const connect = "root:root@tcp(127.0.0.1:3306)/laravel?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./views/html/*.html")
	r.Static("/css", "./views/css")

	r.GET("/", index)
	r.GET("/detail/:id", detail)
	r.GET("/delete/:id", delete)
	r.GET("/update/:id", update)

	r.POST("/do_delete/:id", do_delete)
	r.POST("/do_update/:id", do_update)

	r.Run(":80")
}
