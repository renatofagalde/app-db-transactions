package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const connStr = "postgres://postgres:postgres@localhost:5432/postgres"

func main() {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	defer sqlDB.Close()
}
