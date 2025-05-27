package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mapel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	JurusanID uuid.UUID `json:"ketua_id" gorm:"not null"`
	Name      string    `json:"nama" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
