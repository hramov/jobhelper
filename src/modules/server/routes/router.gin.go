package routes

import (
	gin "github.com/gin-gonic/gin"
	"github.com/hramov/jobhelper/src/modules/server/controller"
)

func Register(router *gin.Engine) {

	deviceController := controller.DeviceController{}
	mainController := controller.MainController{}

	main := router.Group("/")
	{
		main.GET("/", mainController.HomePage)
	}

	device := router.Group("/device")
	{
		device.GET("/", deviceController.FindAll)
	}
}
