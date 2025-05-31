package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ruangan struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string    `json:"nama" gorm:"not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}
