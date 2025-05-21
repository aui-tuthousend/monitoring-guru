package ketua

import (
	e "monitoring-guru/entities"

	"gorm.io/gorm"
)

func CreateKetuaKelas(db *gorm.DB, ketua *e.KetuaKelas) error {
	return db.Create(ketua).Error
}

func GetKetuaKelas(db *gorm.DB, id string) (*e.KetuaKelas, error) {
	var ketua e.KetuaKelas
	if err := db.Where("id = ?", id).First(&ketua).Error; err != nil {
		return nil, err
	}
	return &ketua, nil
}

func GetKetuaKelasByNISN(db *gorm.DB, nip string) (*e.KetuaKelas, error) {
	var ketua e.KetuaKelas
	if err := db.Where("nisn = ?", nip).First(&ketua).Error; err != nil {
		return nil, err
	}
	return &ketua, nil
}

func GetAllKetuaKelas(db *gorm.DB) ([]e.KetuaKelas, error) {
	var ketuaKelas []e.KetuaKelas
	if err := db.Find(&ketuaKelas).Error; err != nil {
		return nil, err
	}
	return ketuaKelas, nil
}

func UpdateKetuaKelas(db *gorm.DB, ketua *e.KetuaKelas) error {
	return db.Save(ketua).Error
}

func DeleteKetuaKelas(db *gorm.DB, id string) error {
	var ketua e.KetuaKelas
	if err := db.Where("id = ?", id).First(&ketua).Error; err != nil {
		return err
	}
	return db.Delete(&ketua).Error
}
