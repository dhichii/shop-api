package database

import (
	"fmt"
	"log"
	"shop-api/src/config/env"
	"shop-api/src/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitMySQL() *gorm.DB {
	config := env.GetSQLEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return db
}

func MigrateMySQL() {
	InitMySQL().AutoMigrate(
		&model.User{},
		&model.Toko{},
		&model.Category{},
		&model.Alamat{},
		&model.Produk{},
		&model.LogProduk{},
		&model.FotoProduk{},
		&model.Trx{},
		&model.DetailTrx{},
	)
}
