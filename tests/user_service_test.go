package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"zhiming.cool/go/internal/routes"
	"zhiming.cool/go/internal/services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userService := services.NewUserService()
	userService.CreateUser("Testing", "Testing@example.com")
	routes.RegisterRoutes(r, userService)
	return r
}

func TestGetUserByName(t *testing.T) {
	router := SetupRouter()

	req, _ := http.NewRequest("GET", "/api/user/Testing", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Testing")
	assert.Contains(t, w.Body.String(), "Testing@example.com")
}

func TestGetUserByNameNotFound(t *testing.T) {
	router := SetupRouter()

	req, _ := http.NewRequest("GET", "/api/user/exists", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")
}

func TestCreateUser(t *testing.T) {
	router := SetupRouter()

	userJSON := `{"name":"John","email":"john@example.com"}`
	req, _ := http.NewRequest("POST", "/api/user", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "John")
	assert.Contains(t, w.Body.String(), "john@example.com")
}

func TestDeleteUser(t *testing.T) {
	router := SetupRouter()

	req, _ := http.NewRequest("DELETE", "/api/user/Testing", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestUpdateUser(t *testing.T) {
	router := SetupRouter()

	userJSON := `{"email":"Testing_new@example.com"}`
	req, _ := http.NewRequest("PUT", "/api/user/Testing", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Testing_new@example.com")
}
