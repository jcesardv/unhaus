package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Unhaus struct {
	ID			uint64	`json:"id" gorm:"primaryKey"`
	Redirect	string	`json:"redirect" gorm:"not null"`
	Unhaus		string	`json:"unhaus" gorm:"unique;not null"`
	Clicked		uint64	`json:"clicked"`
	Random		bool	`json:"random"`
}

func Setup() {
	dsn := "postgresql://dev:NoBA6Lw12dzj7KK-d8xykQ@free-tier14.aws-us-east-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&options=--cluster%3Ddual-spider-1965"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Unhaus{})
	if err != nil {
		fmt.Println(err)
	}
}