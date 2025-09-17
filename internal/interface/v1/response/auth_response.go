package response

type AuthResponse struct {
	Token     string       `json:"token"`
	TokenType string       `json:"token_type"`
	ExpiresIn int          `json:"expires_in"` // seconds
	User      UserResponse `json:"user"`
}

type UserResponse struct {
	ID          uint   `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	IsConfirmed bool   `json:"is_confirmed"`
}
