package routes

import(
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.POST("/users", controller.GetUser())
	incomingRoutes.GET("/users/:id", controller.SignUp())
	incomingRoutes.POST("/users/:id", controller.Login())
}