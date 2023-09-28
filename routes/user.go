package routes

import (
	"github.com/Bhuwanadhikari346/go-gin-crud/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine){
	router.GET("/user", controller.UserController)
}