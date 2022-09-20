package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() *gorm.DB {
	// Внутри этой функции создаём новое соединение методом gorm.Open,
	//  где указываем какой тип базы данных мы хотим использовать и как получить к ней доступ.
	//  В настоящее время Gorm поддерживает только четыре типа баз данных SQL

	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=musicstore password=dev sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Track{})

	return db
}
