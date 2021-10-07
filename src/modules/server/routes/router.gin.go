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

	// auth := router.Group("/auth")
	// {
	// 	auth.GET("/users", userController.Get)
	// 	auth.POST("/login", userController.Login)
	// 	auth.POST("/register", userController.Register)
	// 	auth.GET("/check-jwt", userController.CheckJwt)
	// }

	// expert := router.Group("/experts")
	// {
	// 	expert.GET("/", userController.GetExperts)
	// 	expert.DELETE("/:id", userController.DeleteExpert)
	// }

	// order := router.Group("/orders")
	// {
	// 	order.GET("/:id", orderController.GetById)
	// 	order.POST("/unauth", orderController.CreateUnauth)
	// 	order.GET("/", orderController.Get)
	// 	order.POST("/", orderController.Create)
	// 	order.GET("/client", orderController.GetByClientId)
	// 	order.PATCH("/expert/:order_id", orderController.GetWork)
	// 	order.PATCH("/upload/:id", orderController.UploadDocs)
	// 	order.PATCH("/delete/:id", orderController.DeleteDocs)
	// 	// order.PATCH("/:id", orderController.AssingToClient)
	// }

	// feedback := router.Group("/feedback")
	// {
	// 	feedback.POST("/", feedbackController.Create)
	// 	feedback.GET("/", feedbackController.Get)
	// 	feedback.GET("/:id", feedbackController.GetById)
	// 	// feedback.DELETE("/:id", feedbackController.Delete)
	// }

	// telegram := router.Group("/telegram")
	// {
	// 	telegram.GET("/auth", telegramController.Auth)
	// 	telegram.GET("/news", telegramController.GetNews)
	// }

	// web := router.Group("/")
	// {
	// 	web.GET("/", renderController.VueRenderer)
	// }
}
