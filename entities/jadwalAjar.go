package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JadwalAjar struct {
	ID        	uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	GuruID    	uuid.UUID      `json:"guru_id" gorm:"not null"`
	MapelID   	uuid.UUID      `json:"mapel_id" gorm:"null"`
	KelasID   	uuid.UUID      `json:"kelas_id" gorm:"null"`
	Hari      	string         `json:"hari" gorm:"not null"`
	JamMulai  	string         `json:"jam_mulai" gorm:"not null"`
	JamSelesai 	string         `json:"jam_selesai" gorm:"not null"`
	LastEditor 	string         `json:"last_editor" gorm:"not null"`
	CreatedAt 	time.Time      `json:"created_at"`
	UpdatedAt 	time.Time      `json:"updated_at"`
	DeletedAt 	gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
