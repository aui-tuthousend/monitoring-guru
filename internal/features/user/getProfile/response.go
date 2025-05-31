package getprofile

type GetUserResponse struct {
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
}