package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KetuaKelas struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string         `json:"nama" gorm:"not null"`
	Nisn      string         `json:"nis" gorm:"not null;unique"`
	Password  string         `json:"password" gorm:"not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggerignore:"true"`
}
