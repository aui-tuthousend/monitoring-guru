package database

import (
    "os"
    "log"

	e "monitoring-guru/entities"
    
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Connect() *gorm.DB {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL is not set")
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    if os.Getenv("ENV") != "production" {
	    db.AutoMigrate(&e.Guru{}, &e.Jurusan{}, &e.KetuaKelas{}, &e.Ruangan{}, &e.Mapel{}, &e.Kelas{}) //only call when in local
    }
    return db
}