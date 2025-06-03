package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AbsenMasuk struct {
	ID        	uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	GuruID    	uuid.UUID      `json:"guru_id" gorm:"not null"`
	JadwalAjarID   	uuid.UUID      `json:"jadwal_ajar_id" gorm:"null"`
	RuanganID   	uuid.UUID      `json:"ruangan_id" gorm:"null"`
	Tanggal      	string         `json:"tanggal" gorm:"not null"`
	JamMasuk  	time.Time         `json:"jam_masuk" gorm:"not null"`
	CreatedAt 	time.Time      `json:"-"`
	UpdatedAt 	time.Time      `json:"-"`
	DeletedAt 	gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Guru 		Guru `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:GuruID;references:ID"`
	JadwalAjar 	JadwalAjar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:JadwalAjarID;references:ID"`
	Ruangan 	Ruangan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RuanganID;references:ID"`
}
