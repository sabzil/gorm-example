package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint `gorm:"default:0"`
}

func main() {
	fmt.Println("데이터 베이스 연결")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("데이터베이스 연결 실패")
	}

	db.AutoMigrate(&Product{})

	db.Close()
	fmt.Println("데이터 베이스 연결 종료")
}
