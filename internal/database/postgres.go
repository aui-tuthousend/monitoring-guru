package database

import (
    "os"
    "log"

	e "monitoring-guru/entities"
	"monitoring-guru/docs"
    
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
		docs.SwaggerInfo.Host = "127.0.0.1:8080"
	    db.AutoMigrate(
            &e.Guru{}, 
            &e.Jurusan{}, 
            &e.KetuaKelas{}, 
            &e.Ruangan{}, 
            &e.Mapel{}, 
            &e.Kelas{}, 
            &e.JadwalAjar{},
            &e.StatusKelas{},
            &e.AbsenMasuk{},
            &e.AbsenKeluar{},
            &e.Izin{},
        ) //only call when in local
    } else {
		docs.SwaggerInfo.Host = "monitoring.aui-tuthousend.cyou" // change later
    }
    return db
}