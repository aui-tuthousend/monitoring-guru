package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AbsenKeluar struct {
	ID        	uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	AbsenMasukID    	uuid.UUID      `json:"absen_masuk_id" gorm:"not null"`
	JamKeluar   	time.Time         `json:"jam_keluar" gorm:"not null"`
	Status   	string         `json:"status"`
	CreatedAt 	time.Time      `json:"created_at"`
	UpdatedAt 	time.Time      `json:"updated_at"`
	DeletedAt 	gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
