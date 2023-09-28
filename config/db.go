package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
func connect(){
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:8080/postgres"),
	if err != nil{
		panic(err)
	}
	db.Automigrate(&gorm.Model.user{})
	DB = db

)}