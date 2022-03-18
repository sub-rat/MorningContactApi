package postgres

import (
	"github.com/sub-rat/MorningContactApi/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=MoringContactApi port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&models.User{},
		&models.Contact{},
		&models.Address{},
		&models.Phone{},
	)
	return db
}
