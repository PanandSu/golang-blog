package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	username := "root"
	password := "(PanandSu521125)"
	host := "127.0.0.1:3306"
	dbname := "exercise"
	other := "charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", username, password, host, dbname, other)

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
	r := gin.Default()
	port := "8083"
	_ = r.Run(":" + port)
}
func main() {

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
