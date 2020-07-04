package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sasa-radovanovic/bookstore_users-api/domain/users"
	"github.com/sasa-radovanovic/bookstore_users-api/services"
	"github.com/sasa-radovanovic/bookstore_users-api/utils/errors"
)

// GetUser handler
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("invalid user id")
		c.JSON(restErr.Code, restErr)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser creates user
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// SearchUser finds the user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not ready yet")
}
