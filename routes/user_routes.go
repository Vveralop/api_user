package routes

import (
	"gogin/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
    userRoutes := router.Group("/users")
    {
        userRoutes.GET("/", controller.GetUsers)
        userRoutes.GET("/:id", controller.GetUserByID)
        userRoutes.POST("/", controller.CreateUser)
        userRoutes.PUT("/:id", controller.UpdateUser)
        userRoutes.DELETE("/:id", controller.DeleteUser)
    }
}
