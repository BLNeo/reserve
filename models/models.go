package models

import (
	"gorm.io/gorm"
	"reserve/tool/log"
	"reserve/tool/mysql"
	"sync"
	"time"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDb(dbEngine *gorm.DB) {
	once.Do(func() {
		db = dbEngine
	})
}

func GetDb() *gorm.DB {
	return db
}

type Model struct {
	ID        int32 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// AutoMigrate 表同步
func AutoMigrate() {
	tables := []interface{}{
		new(Book),
	}
	err := mysql.AutoMigrate(db, tables...)
	if err != nil {
		log.Logger.Error("mysql AutoMigrate err:" + err.Error())
		panic(err)
	}
}
