package users

import "encoding/json"

// PublicUser available publicly
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// PrivateUser accessed privately
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// Marshall maps user to appropriate class
func (user *User) Marshall(isPublic bool) interface{} {
	userJSON, _ := json.Marshal(user)
	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(userJSON, &publicUser)
		return publicUser
	}

	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}

// Marshall maps user to appropriate class
func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}
