package jadwalajar

import (
	e "monitoring-guru/entities"

	"gorm.io/gorm"
)

func CreateJadwalajar(db *gorm.DB, jadwalajar *e.JadwalAjar) error {
	return db.Create(jadwalajar).Error
}

func UpdateJadwalajar(db *gorm.DB, jadwalajar *e.JadwalAjar) error {
	return db.Save(jadwalajar).Error
}

func GetJadwalajarByID(db *gorm.DB, id string) (*e.JadwalAjar, error) {
	var jadwalajar e.JadwalAjar
	if err := db.Where("id = ?", id).First(&jadwalajar).Error; err != nil {
		return nil, err
	}
	return &jadwalajar, nil
}

func GetAllJadwalajar(db *gorm.DB) ([]e.JadwalAjar, error) {
	var jadwalajarList []e.JadwalAjar
	if err := db.
		Preload("Mapel").
		Preload("Guru").
		Preload("Kelas").
		Find(&jadwalajarList).Error; err != nil {
		return nil, err
	}
	return jadwalajarList, nil
}


func GetJadwalajarByIDGuru(db *gorm.DB, id string, hari string) ([]e.JadwalAjar, error) {
	var jadwalajarList []e.JadwalAjar

	if err := db.
		Preload("Mapel").
		// Preload("Guru").
		Preload("Kelas").
		Where("guru_id = ? AND hari = ?", id, hari).
		Find(&jadwalajarList).Error; err != nil {
		return nil, err
	}

	return jadwalajarList, nil
}


func GetJadwalajarByIDKelas(db *gorm.DB, id string, hari string) ([]e.JadwalAjar, error) {
	var jadwalajarList []e.JadwalAjar

	if err := db.
		Preload("Mapel").
		Preload("Guru").
		// Preload("Kelas").
		Where("kelas_id = ? AND hari = ?", id, hari).
		Find(&jadwalajarList).Error; err != nil {
		return nil, err
	}

	return jadwalajarList, nil
}


