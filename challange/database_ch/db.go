package database_ch

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"challange/m/modules_ch"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDb() {
	dns := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dns))
	if err != nil {
		log.Fatal("Error connect database", err)
	}
	DB.AutoMigrate(modules_ch.Revenue{})
	DB.AutoMigrate(modules_ch.Expense{})
}
