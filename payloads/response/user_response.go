package response

import "microservices/models"

// UserResponse struct for response of user
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

// Transform from models.User to UserResponse
func (u *UserResponse) Transform(user models.User) {
	u.ID = user.ID
	u.Username = user.Username
	u.Email = user.Email
	u.IsActive = user.IsActive
}
