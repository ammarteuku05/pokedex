package config

import (
	"fmt"

	"pokedex/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", GetConfig().DbUser, GetConfig().DbPass, GetConfig().DbHost, GetConfig().DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(entity.User{}, entity.Type{}, entity.Monster{}, entity.TypeOfMonster{})
	return db
}
