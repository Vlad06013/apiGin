package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=smm_planer dbname=smm_planer password=eRdHs3gFsa3gt!n sslmode=disable")
	//db, err := gorm.Open("postgres", "host=212.67.8.221 port=5432 user=smm_planer dbname=smm_planer password=eRdHs3gFsa3gt!n sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Bot{})

	return db
}
