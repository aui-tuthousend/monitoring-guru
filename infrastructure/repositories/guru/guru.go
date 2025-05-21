package guru

import (
	e "monitoring-guru/entities"

	"gorm.io/gorm"
)

func CreateGuru(db *gorm.DB, guru *e.Guru) error {
	return db.Create(guru).Error
}

func GetGuru(db *gorm.DB, id string) (*e.Guru, error) {
	var guru e.Guru
	if err := db.Where("id = ?", id).First(&guru).Error; err != nil {
		return nil, err
	}
	return &guru, nil
}

func GetGuruByNIP(db *gorm.DB, nip string) (*e.Guru, error) {
	var guru e.Guru
	if err := db.Where("nip = ?", nip).First(&guru).Error; err != nil {
		return nil, err
	}
	return &guru, nil
}

func GetAllGuru(db *gorm.DB) ([]e.Guru, error) {
	var gurus []e.Guru
	if err := db.Find(&gurus).Error; err != nil {
		return nil, err
	}
	return gurus, nil
}

func UpdateGuru(db *gorm.DB, guru *e.Guru) error {
	return db.Save(guru).Error

}
func DeleteGuru(db *gorm.DB, id string) error {
	var guru e.Guru
	if err := db.Where("id = ?", id).First(&guru).Error; err != nil {
		return err
	}
	return db.Delete(&guru).Error
}
