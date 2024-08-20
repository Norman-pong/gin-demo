package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zhiming.cool/go/internal/models"
	"zhiming.cool/go/internal/services"
)

func GetUser(c *gin.Context, userService *services.UserService) {
	name := c.Param("name")
	user, err := userService.GetUserByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context, userService *services.UserService) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := userService.CreateUser(user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

func DeleteUser(c *gin.Context, userService *services.UserService) {
	name := c.Param("name")
	if err := userService.DeleteUser(name); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func UpdateUser(c *gin.Context, userService *services.UserService) {
	name := c.Param("name")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := userService.UpdateUser(name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}
