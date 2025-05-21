package auth

import (
	e "monitoring-guru/entities"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateJWT(userID uuid.UUID, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func FindUserByEmail(db *gorm.DB, email string) (*e.User, error) {
	var user e.User
	result := db.Raw("SELECT * FROM users WHERE email = ? LIMIT 1", email).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}

func FindGuruByNip(db *gorm.DB, nip string) (*e.Guru, error) {
	var guru e.Guru
	result := db.Raw("SELECT * FROM gurus WHERE nip = ? LIMIT 1", nip).Scan(&guru)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &guru, nil
}

func FindKetuaKelasByNisn(db *gorm.DB, nisn string) (*e.KetuaKelas, error) {
	var ketuaKelas e.KetuaKelas
	result := db.Raw("SELECT * FROM ketuakelas WHERE nisn = ? LIMIT 1", nisn).Scan(&ketuaKelas)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &ketuaKelas, nil
}
