package routes

import (
	"github.com/gin-gonic/gin"
	"zhiming.cool/go/internal/handlers"
	"zhiming.cool/go/internal/services"
)

func RegisterRoutes(r *gin.Engine, userService *services.UserService) {
	api := r.Group("/api")
	{
		api.GET("/user", func(c *gin.Context) {
			handlers.GetUsers(c, userService)
		})
		api.GET("/user/:id", func(c *gin.Context) {
			handlers.GetUser(c, userService)
		})
		api.PUT("/user/:id", func(c *gin.Context) {
			handlers.UpdateUser(c, userService)
		})
		api.POST("/user", func(c *gin.Context) {
			handlers.CreateUser(c, userService)
		})
		api.DELETE("/user/:id", func(c *gin.Context) {
			handlers.DeleteUser(c, userService)
		})
	}
}
