package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JadwalAjar struct {
	ID        	uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	GuruID    	uuid.UUID      `json:"guru_id" gorm:"not null"`
	Guru 		Guru `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:GuruID;references:ID"`
	MapelID   	uuid.UUID      `json:"mapel_id" gorm:"null"`
	Mapel 		Mapel	 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:MapelID;references:ID"`
	KelasID   	uuid.UUID      `json:"kelas_id" gorm:"null"`
	Kelas 		Kelas `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:KelasID;references:ID"`
	Hari      	string         `json:"hari" gorm:"not null"`
	JamMulai  	string         `json:"jam_mulai" gorm:"not null"`
	JamSelesai 	string         `json:"jam_selesai" gorm:"not null"`
	LastEditor 	string         `json:"last_editor" gorm:"not null"`
	CreatedAt 	time.Time      `json:"created_at"`
	UpdatedAt 	time.Time      `json:"updated_at"`
	DeletedAt 	gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
