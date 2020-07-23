package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sasa-radovanovic/bookstore_users-api/domain/users"
	"github.com/sasa-radovanovic/bookstore_users-api/services"
	"github.com/sasa-radovanovic/bookstore_users-api/utils/errors"
)

const (
	userIDPathParam = "user_id"
)

// GetUser handler
func GetUser(c *gin.Context) {
	userID, userErr := getUserIDPathParam(c)
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}
	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

// CreateUser creates user
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// UpdateUser updates given user
func UpdateUser(c *gin.Context) {
	userID, userErr := getUserIDPathParam(c)
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}
	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch
	result, saveErr := services.UsersService.UpdateUser(user, isPartial)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// DeleteUser deletes user from the database
func DeleteUser(c *gin.Context) {
	userID, userErr := getUserIDPathParam(c)
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}
	result, saveErr := services.UsersService.DeleteUser(userID)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Extracts user_id param from URL path
func getUserIDPathParam(c *gin.Context) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}
	return userID, nil
}

// FindByStatus executes a search based on status
func FindByStatus(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

// Login to login user
func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}
	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
