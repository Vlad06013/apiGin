package routes

import (
	"github.com/Vlad06013/apiGin/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	router.GET("/telegram-input", controllers.InputTGRequest)

}
