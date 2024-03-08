package router

import (
	"github.com/gin-gonic/gin"
	"golang-blog/controller"
)

func Routers(router *gin.Engine) {
	root := router.Group("/")
	fromRoot(root)
	_ = router.Run("8083")
}

func fromRoot(group *gin.RouterGroup) {
	group.GET("/", controller.IndexPage)
	group.GET("register", controller.Register)
	group.GET("login", controller.Login)
	group.GET("logout", controller.Logout)
	group.GET("posts", controller.GetAllPosts)
	post := group.Group("post")
	fromPost(post)
	group.GET("user/posts/:username", controller.GetUserPost)
}

func fromPost(group *gin.RouterGroup) {
	group.GET("/:id", controller.GetPost)
	group.GET("/create", controller.CreatePost)
	group.GET("/delete/:id", controller.DeletePost)
	group.GET("/edit/:id", controller.EditPost)
}
