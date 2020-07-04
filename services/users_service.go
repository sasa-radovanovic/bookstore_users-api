package services

import (
	"github.com/sasa-radovanovic/bookstore_users-api/domain/users"
	"github.com/sasa-radovanovic/bookstore_users-api/utils/errors"
)

// CreateUser - Create and persist user in DB
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser retrieves user
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := users.User{
		ID: userID,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}
