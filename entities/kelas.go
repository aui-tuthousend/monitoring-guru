package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kelas struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	KetuaID   uuid.UUID      `json:"ketua_id" gorm:"not null"`
	WakilID   uuid.UUID      `json:"wakil_id" gorm:"null"`
	JurusanID uuid.UUID      `json:"jurusan_id" gorm:"null"`
	Nama      string         `json:"nama" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
