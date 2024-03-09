package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	//c.JSON(200, "Index Page")
}

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.tmpl", gin.H{})
	//c.JSON(200, "Register Page")
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	//c.JSON(200, "Login Page")
}

func Logout(c *gin.Context) {
	c.HTML(http.StatusOK, "logout.tmpl", gin.H{})
	//c.JSON(200, "Logout Page")
}

func GetAllPosts(c *gin.Context) {
	c.HTML(http.StatusOK, "posts.tmpl", gin.H{})
	//c.JSON(200, "Get All Posts")
}

func GetPost(c *gin.Context) {
	c.HTML(http.StatusOK, "post.tmpl", gin.H{})
	//c.JSON(200, "Get Post")
}

func CreatePost(c *gin.Context) {
	c.HTML(http.StatusOK, "create.tmpl", gin.H{})
	//c.JSON(200, "Create Post")
}

func EditPost(c *gin.Context) {
	c.HTML(http.StatusOK, "edit.tmpl", gin.H{})
	//c.JSON(200, "Edit Post")
}

func DeletePost(c *gin.Context) {
	c.HTML(http.StatusOK, "delete.tmpl", gin.H{})
	//c.JSON(200, "Delete Post")
}

func GetUserPost(c *gin.Context) {
	c.HTML(http.StatusOK, "user_post.tmpl", gin.H{})
	//c.JSON(200, "Get User Post")
}
