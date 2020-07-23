package users

// LoginRequest is DTO for logging in
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
