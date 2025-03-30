package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Database() *gorm.DB {
	dsn := os.Getenv("DSN")
	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic("failed to connect database")
	} else {
		return db
	}
}
