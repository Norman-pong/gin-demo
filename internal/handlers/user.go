package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zhiming.cool/go/internal/models"
	"zhiming.cool/go/internal/services"
)

func GetUser(c *gin.Context, userService *services.UserService) {
	cid := c.Param("id")
	user, err := userService.GetUserById(cid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context, userService *services.UserService) {
	users, err := userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(users) == 0 {
		c.JSON(http.StatusNoContent, []models.User{})
		return
	}
	c.JSON(http.StatusOK, users)
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
	cid := c.Param("id")
	if _, err := userService.DeleteUser(cid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func UpdateUser(c *gin.Context, userService *services.UserService) {
	cid := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := userService.UpdateUser(cid, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}
