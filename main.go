package main

import (
	"github.com/gin-gonic/gin"
	"golang-blog/router"
)

func main() {
	r := gin.Default()
	router.Routers(r)
}
