package mysql

import (
	"todo/infra/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Connect 接続
func Connect() *gorm.DB {
	var err error
	db, err = gorm.Open("mysql", "root:mysql@tcp(mysql:3306)/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// Migrate マイグレーション
func Migrate() {
	db.AutoMigrate(&model.Todo{})
}

// CloseConnect 接続を切る
func CloseConnect() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func DefaultDB() *gorm.DB {
	return db
}
