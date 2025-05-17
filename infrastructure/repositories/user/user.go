package user

import (
	e "monitoring-guru/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func GetUser(db *gorm.DB, id string) (*e.User, error) {
    var user e.User
    result := db.Raw("SELECT * FROM users WHERE id = ? LIMIT 1", id).Scan(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func CreateUser(db *gorm.DB, user *e.User) error {
    return db.Create(user).Error
}