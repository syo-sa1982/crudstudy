package main

import (

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/syo-sa1982/crudstudy/models"
)

func main() {
	db, _ := gorm.Open("mysql", "root:@/gorm?charset=utf8&parseTime=True")
	db.CreateTable(&models.User{})
}
