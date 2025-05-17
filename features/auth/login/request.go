package login


// LoginRequestBody
// @Description Login user request body
type AuthRequest struct {
    // Your Email
    Email    string `json:"email"`
    // Your Password 
    Password string `json:"password"`
}
