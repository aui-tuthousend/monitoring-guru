package create

type CreateUserResponse struct {
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
}

type CreateUserResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    CreateUserResponse  `json:"data"`
}
