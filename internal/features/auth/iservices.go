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

type AuthService struct {
	DB *gorm.DB
}

func (a *AuthService) CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (a *AuthService) GenerateJWT(userID uuid.UUID, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID.String(),
		"role": role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (a *AuthService) FindUserByEmail(email string) (*e.User, error) {
	var user e.User
	result := a.DB.Raw("SELECT * FROM users WHERE email = ? LIMIT 1", email).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}

func (a *AuthService) FindGuruByNip(nip string) (*e.Guru, error) {
	var guru e.Guru
	result := a.DB.Raw("SELECT * FROM gurus WHERE nip = ? LIMIT 1", nip).Scan(&guru)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &guru, nil
}

func (a *AuthService) FindKetuaKelasByNisn(nisn string) (*e.KetuaKelas, error) {
	var ketuaKelas e.KetuaKelas
	result := a.DB.Raw("SELECT * FROM ketua_kelas WHERE nisn = ? LIMIT 1", nisn).Scan(&ketuaKelas)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &ketuaKelas, nil
}