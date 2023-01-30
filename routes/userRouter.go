package routes

import (
	controller "github.com/MikeMwita/JWT_project/controllers"
	"github.com/MikeMwita/JWT_project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate)
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:users_id", controller.GetUser())

}
