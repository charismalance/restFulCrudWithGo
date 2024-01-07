package router

import (
	userRouter "alireza/controller"
	"github.com/gin-gonic/gin"
)

func Router (){
	router := gin.Default()

	router.POST("/findUserByName",userRouter.FindUserByName)
	router.POST("/findUserById",userRouter.FindUserById)
	router.POST("/findAllUser",userRouter.FindAllUser)
	router.Run()
}