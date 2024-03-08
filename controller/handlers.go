package controller

import "github.com/gin-gonic/gin"

func IndexPage(c *gin.Context) {
	c.JSON(200, "Index Page")
}

func Register(c *gin.Context) {
	c.JSON(200, "Register Page")
}

func Login(c *gin.Context) {
	c.JSON(200, "Login Page")
}

func Logout(c *gin.Context) {
	c.JSON(200, "Logout Page")
}

func GetAllPosts(c *gin.Context) {
	c.JSON(200, "Get All Posts")
}

func GetPost(c *gin.Context) {
	c.JSON(200, "Get Post")
}

func CreatePost(c *gin.Context) {
	c.JSON(200, "Create Post")
}

func EditPost(c *gin.Context) {
	c.JSON(200, "Edit Post")
}

func DeletePost(c *gin.Context) {
	c.JSON(200, "Delete Post")
}

func GetUserPost(c *gin.Context) {
	c.JSON(200, "Get User Post")
}
