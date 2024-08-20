package tests

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"zhiming.cool/go/internal/routes"
	"zhiming.cool/go/internal/services"
)

func SetupRouter() (*gin.Engine, *services.UserService) {
	userService := services.NewUserService("file::memory:?cache=shared&")
	r := gin.Default()
	routes.RegisterRoutes(r, userService)
	return r, userService
}

func TestGetUserByNameNotFound(t *testing.T) {
	router, _ := SetupRouter()

	req, _ := http.NewRequest("GET", "/api/user/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")
}

func TestGetUsersEmpty(t *testing.T) {
	router, _ := SetupRouter()

	req, _ := http.NewRequest("GET", "/api/user", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Contains(t, w.Body.String(), "")
}

func TestGetUserById(t *testing.T) {
	router, userService := SetupRouter()

	userService.CreateUser("Testing", "Testing@example.com")

	req, _ := http.NewRequest("GET", "/api/user/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Testing")
	assert.Contains(t, w.Body.String(), "Testing@example.com")
}

func TestCreateUser(t *testing.T) {
	router, _ := SetupRouter()

	userJSON := `{"name":"John","email":"john@example.com"}`
	req, _ := http.NewRequest("POST", "/api/user", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "John")
	assert.Contains(t, w.Body.String(), "john@example.com")
}

func TestGetUsers(t *testing.T) {
	router, userService := SetupRouter()

	userService.CreateUser("Testing", "Testing@example.com")
	userService.CreateUser("Testing2", "Testing2@example.com")

	req, _ := http.NewRequest("GET", "/api/user", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Testing")
	assert.Contains(t, w.Body.String(), "Testing2")
}

func TestUpdateUser(t *testing.T) {
	router, _ := SetupRouter()

	userJSON := `{"email":"Testing_new@example.com"}`
	req, _ := http.NewRequest("PUT", "/api/user/1", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Testing_new@example.com")
}

func TestDeleteUser(t *testing.T) {
	router, userService := SetupRouter()

	var req *http.Request

	req, _ = http.NewRequest("DELETE", "/api/user/100", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	user, _ := userService.CreateUser("TestingDelete", "Testing@mail.com")
	req, _ = http.NewRequest("DELETE", "/api/user/"+strconv.Itoa(int(user.ID)), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
