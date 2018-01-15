package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code     string
	TargetID string
	Price    uint `gorm:"default:0"`
}

func main() {
	fmt.Println("데이터 베이스 연결")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("데이터베이스 연결 실패")
	}

	db.AutoMigrate(&Product{})

	p := &Product{Code: "1111", TargetID: "ABC", Price: 1000}

	tx := db.Begin()

	tmp := &Product{}
	if tx.Where("Code = 1111").First(&tmp).RecordNotFound() == true {
		fmt.Println("Record Not Found")
		if tx.Create(p).Error != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	} else {
		fmt.Println(tmp)
	}

	db.Close()
	fmt.Println("데이터 베이스 연결 종료")
}
