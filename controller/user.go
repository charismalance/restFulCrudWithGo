package userRouter

import (
	"alireza/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const dsn = " root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
var db,errors = gorm.Open(mysql.Open(dsn), &gorm.Config{})


func FindUserByName(ctx *gin.Context) {
	var users []model.User
	fmt.Println(dsn)
	if errors != nil {
		log.Println(4656654, errors)
	}
	var p model.Body
	if err := ctx.ShouldBindJSON(&p); err != nil {
		fmt.Println(err)
	}
	db.Where("name =?", p.Name).Find(&users)
	ctx.JSON(200, users)

}
func FindUserById(c *gin.Context){
	var users []model.User 
	var p model.Body
	 if err :=c.ShouldBindJSON(&p);err != nil {
		panic(err)
	 }
	 if errors != nil {panic(errors)}
	 db.Where("id = ?" , p.ID).Find(&users)
	 c.JSON(200, users)

}

func FindAllUser(c *gin.Context){
	var users []model.User
	var find model.Body
	if err := c.ShouldBindJSON(&find) ; err != nil {
		panic(err)
	}
	if errors != nil {panic(errors)}
	db.Find(&users)
	c.JSON(200,users)
}
