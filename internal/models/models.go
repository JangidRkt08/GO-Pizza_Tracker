package models

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type DBModel struct {
	Order OrderModel
}

func InitDB(dataSourceName string) (*DBModel, error){
	db , err := gorm.Open(sqlite.Open(dataSourceName),&gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate Db %v", err)

		
	}
	err = db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate Db %v", err)
	}
	dbModel := &DBModel{
		Order :OrderModel{DB: db},
	}

	return dbModel, nil
}