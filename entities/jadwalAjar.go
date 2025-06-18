package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JadwalAjar struct {
	ID        	uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	GuruID    	uuid.UUID      `json:"guru_id" gorm:"not null; type:uuid"`
	MapelID   	uuid.UUID      `json:"mapel_id" gorm:"null; type:uuid"`
	KelasID   	uuid.UUID      `json:"kelas_id" gorm:"null; type:uuid"`
	RuanganID   	uuid.UUID      `json:"ruangan_id" gorm:"null; type:uuid"`
	Hari     	string         `json:"hari" gorm:"not null"`
	JamMulai  	string        `json:"jam_mulai" gorm:"not null"`
	JamSelesai 	string         `json:"jam_selesai" gorm:"not null"`
	LastEditor 	string         `json:"last_editor"`
	CreatedAt 	time.Time      `json:"-"`
	UpdatedAt 	time.Time      `json:"-"`
	DeletedAt 	gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggerignore:"true"`

	Guru 		Guru `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:GuruID;references:ID"`
	Mapel 		Mapel	 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:MapelID;references:ID"`
	Kelas 		Kelas `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:KelasID;references:ID"`
	Ruangan 	Ruangan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RuanganID;references:ID"`
}
