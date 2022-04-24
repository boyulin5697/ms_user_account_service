//db connection initialize, Author @by. since 2022/4/22

package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitDB() *gorm.DB {
	destination := "ms_user_account:a7782114312@tcp(106.55.179.158:3306)/ms_user_account?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(destination), &gorm.Config{})
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println(err)
	return db
}
