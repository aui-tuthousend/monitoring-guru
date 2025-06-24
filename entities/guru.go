package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Guru struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string         `json:"name" gorm:"not null"`
	Nip       string         `json:"nip" gorm:"not null"`
	Jabatan   string         `json:"jabatan" gorm:"not null"`
	Password  string         `json:"password" gorm:"not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggerignore:"true"`
}
