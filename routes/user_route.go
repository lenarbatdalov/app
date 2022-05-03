package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lenarbatdalov/app/controllers"
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", controllers.GetAllUsers())
	r.POST("/user", controllers.CreateUser())
	r.GET("/user/:userId", controllers.GetAUser())
	r.PUT("/user/:userId", controllers.EditAUser())
	r.DELETE("/user/:userId", controllers.DeleteAUser())
}
