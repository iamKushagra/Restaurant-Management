package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.POST("/users", controller.GetUser())
	incomingRoutes.GET("/users/:id", controller.SignUp())
	incomingRoutes.POST("/users/:id", controller.Login())
}
