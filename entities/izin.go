package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Izin struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	GuruID       uuid.UUID      `json:"guru_id" gorm:"type:uuid;not null"`
	JadwalAjarID uuid.UUID      `json:"jadwal_ajar_id" gorm:"type:uuid;not null"`
	TanggalIzin  time.Time      `json:"tanggal_izin" gorm:"not null"`
	Pesan        string         `json:"pesan" gorm:"type:text"`
	Approval     bool           `json:"approval" gorm:"default:false"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggerignore:"true"`

	Guru       Guru       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:GuruID;references:ID"`
	JadwalAjar JadwalAjar `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:JadwalAjarID;references:ID"`
}
