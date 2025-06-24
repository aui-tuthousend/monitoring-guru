package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusKelas struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	KelasID   uuid.UUID      `json:"kelas_id" gorm:"not null;type:uuid"`
	Kelas	  Kelas	 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:KelasID;references:ID"`
	Mapel string `json:"mapel" gorm:"null"`
	Pengajar string `json:"pengajar" gorm:"null"`
	Ruangan string `json:"ruangan" gorm:"null"`
	IsActive  bool           `json:"is_active" gorm:"default:false"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggerignore:"true"`
}
