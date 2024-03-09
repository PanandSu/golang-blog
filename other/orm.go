package other

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
)

func orm() {
	username := "root"
	password := "(PanandSu521125)"
	host := "127.0.0.1:3306"
	dbname := "exercise"
	other := "charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		username,
		password,
		host,
		dbname,
		other,
	)

	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		panic(err)
	}

	type List struct {
		gorm.Model
		Name  string `gorm:"type:varchar(255);not null" json:"name" binding:"required"`
		Addr  string `gorm:"type:varchar(255);not null" json:"addr"`
		Phone string `gorm:"type:varchar(255);not null" json:"phone" binding:"required"`
		Age   int    `gorm:"type:int(11);not null" json:"age"`
	}

	err = db.AutoMigrate(&List{})
	if err != nil {
		return
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/user/add", func(c *gin.Context) {
		var user List
		err = c.ShouldBindJSON(user)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  err.Error(),
				"data": nil,
			})
		} else {
			db.Create(&user)
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "ok",
				"data": user,
			})
		}
	})

	r.DELETE("/user/delete/:id", func(c *gin.Context) {
		var users []List
		id := c.Param("id")
		db.Where("id =?", id).Find(&users)
		if len(users) == 0 {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "id 没有找到,删除失败",
			})
		} else {
			db.Where("id = ?", id).Delete(&users)
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "删除成功",
			})
		}
	})

	r.POST("/user/update/:id", func(c *gin.Context) {
		var user List
		id := c.Param("id")
		db.Where("id = ?", id).Find(&user)
		if user.ID == 0 {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "用户id没有找到",
			})
		} else {
			err = c.ShouldBindJSON(&user)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 400,
					"msg":  "修改失败",
				})
			} else {
				db.Where("id = ?", id).Updates(&user)
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "修改成功",
					"data": user,
				})
			}
		}
	})

	r.GET("/user/list/:id", func(c *gin.Context) {
		var users []List
		id := c.Param("id")
		db.Where("id =?", id).Find(&users)
		if len(users) == 0 {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "没有找到该用户",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "ok",
				"data": users,
			})
		}
	})

	r.GET("/user/list", func(c *gin.Context) {
		var users []List
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		pageNum, _ := strconv.Atoi(c.Query("pageNum"))
		if pageSize == 0 {
			pageSize--
		}
		if pageNum == 0 {
			pageNum--
		}
		offset := (pageNum - 1) * pageSize
		if pageSize == -1 && pageNum == -1 {
			offset = -1
		}
		db.Model(&users).Limit(pageSize).Offset(offset).Find(&users)
		if len(users) == 0 {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "没有查询到数据",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "查询成功",
				"data": gin.H{
					"list":     users,
					"pageNum":  pageNum,
					"pageSize": pageSize,
					"total":    len(users),
				},
			})
		}
	})

	port := "8083"
	_ = r.Run(":" + port)
}

/*
蓝色是修改过的,绿色是新建的
package main

import (
	"github.com/gin-gonic/gin"
	"golang-blog/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	router.Routers(r)
}
*/
