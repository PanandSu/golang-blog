package main

import (
	"github.com/gin-gonic/gin"
	"golang-blog/router"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	router.Route(r)
	_ = r.Run(":8083")
}
