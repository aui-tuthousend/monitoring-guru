package create

// CreateUserRequestBody
// @Description Create user request body
type CreateUserRequest struct {
    // Your Name
    Name     string `json:"name"`
    // Your Email
    Email    string `json:"email" gorm:"unique"`
    // Your Password 
    Password string `json:"password"`
}