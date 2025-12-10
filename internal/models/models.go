package models

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type DBModel struct {
	Order OrderModel
	User UserModel
	DB *gorm.DB
}

func InitDB(dataSourceName string) (*DBModel, error){
	db , err := gorm.Open(sqlite.Open(dataSourceName),&gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate Db %v", err)

		
	}
	err = db.AutoMigrate(&Order{}, &OrderItem{}, &User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate Db %v", err)
	}
	dbModel := &DBModel{
		DB : db,
		Order :OrderModel{DB: db},
		User : UserModel{DB: db},
	}

	return dbModel, nil
}