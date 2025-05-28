package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kelas struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	KetuaID   uuid.UUID      `json:"ketua_id" gorm:"not null"`
	Ketua	  KetuaKelas	 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:KetuaID;references:ID"`
	WakilID   uuid.UUID      `json:"wakil_id" gorm:"null"`
	Wakil	  KetuaKelas	 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:WakilID;references:ID"`
	JurusanID uuid.UUID      `json:"jurusan_id" gorm:"null"`
	Jurusan   Jurusan        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:JurusanID;references:ID"`
	Nama      string         `json:"nama" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"default:false"`
	JadwalAjar	     []JadwalAjar   `gorm:"foreignKey:KelasID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
