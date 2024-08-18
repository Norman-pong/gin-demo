package routes

import (
	"github.com/gin-gonic/gin"
	"zhiming.cool/go/internal/handlers"
	"zhiming.cool/go/internal/services"
)

func RegisterRoutes(r *gin.Engine, userService *services.UserService) {
	api := r.Group("/api")
	{
		api.GET("/user/:name", func(c *gin.Context) {
			handlers.GetUser(c, userService)
		})
		api.PUT("/user/:name", func(c *gin.Context) {
			handlers.UpdateUser(c, userService)
		})
		api.POST("/user", func(c *gin.Context) {
			handlers.CreateUser(c, userService)
		})
		api.DELETE("/user/:name", func(c *gin.Context) {
			handlers.DeleteUser(c, userService)
		})
	}
}
