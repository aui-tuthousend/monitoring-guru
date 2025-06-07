package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kelas struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	KetuaID   uuid.UUID      `json:"ketua_id" gorm:"not null;type:uuid"`
	Ketua	  KetuaKelas	 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:KetuaID;references:ID"`
	// WakilID   uuid.UUID      `json:"wakil_id" gorm:"null;type:uuid"`
	// Wakil	  KetuaKelas	 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:WakilID;references:ID"`
	JurusanID uuid.UUID      `json:"jurusan_id" gorm:"null;type:uuid"`
	Jurusan   Jurusan        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:JurusanID;references:ID"`
	Name      string         `json:"name" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"default:false"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggerignore:"true"`
}
