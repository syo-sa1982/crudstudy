package main

import (

	"github.com/jinzhu/gorm"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db gorm.DB

func main() {
	log.Print("main")
	user := web.New()
	goji.Handle("/user/*", user)

	user.Use(middleware.SubRouter)
	user.Use(SuperSecure)
	user.Get("/index", UserIndex)
	user.Get("/new", UserNew)
	user.Post("/new", UserCreate)
	user.Get("/edit/:id", UserEdit)
	user.Post("/update/:id", UserUpdate)
	user.Get("/delete/:id", UserDelete)

	goji.Serve()
}

func init() {
	db, _ = gorm.Open("mysql", "root:@/gorm?charset=utf8&parseTime=True")
}