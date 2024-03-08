package router

import "github.com/gin-gonic/gin"

func Routers(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Index Page")
	})
	_ = router.Run("8083")
}
