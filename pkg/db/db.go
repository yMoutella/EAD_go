package db

import (
	"ymoutella/ead/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func (db *Db) Connect() *Db {
	dsn := "host=localhost user=postgres password=admin dbname=ead_go port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	inDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error trying to connect in database!")
	}

	db.DB = inDb

	db.DB.AutoMigrate(&models.Course{}, &models.Module{}, &models.Lesson{})

	return db
}
