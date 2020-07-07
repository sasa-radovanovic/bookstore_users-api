package services

import (
	"fmt"

	"github.com/sasa-radovanovic/bookstore_users-api/domain/users"
	cryptoutils "github.com/sasa-radovanovic/bookstore_users-api/utils/crypto_utils"
	"github.com/sasa-radovanovic/bookstore_users-api/utils/errors"
)

var (
	// UsersService interface
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(users.User, bool) (*users.User, *errors.RestErr)
	SearchUser(string) (users.Users, *errors.RestErr)
	DeleteUser(int64) (*users.User, *errors.RestErr)
}

// CreateUser - Create and persist user in DB
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.Password = cryptoutils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser retrieves user
func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := users.User{
		ID: userID,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates user
func (s *usersService) UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		fmt.Println(err.Error)
		return nil, err
	}
	return current, nil
}

// DeleteUser deletes user from the database
func (s *usersService) DeleteUser(userID int64) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(userID)
	if err != nil {
		return nil, err
	}
	if err := current.Delete(); err != nil {
		return nil, err
	}
	return current, nil
}

// Search executes a search
func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := users.User{}
	return dao.FindByStatus(status)
}
