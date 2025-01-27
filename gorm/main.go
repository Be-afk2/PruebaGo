package main

import (
	"github.com/culturadevops/GORM/libs"
)

func main() {
	dbConfig := libs.Configure("./", "mysql")
	libs.DB = dbConfig.InitMysqlDB()
	libs.DB.AutoMigrate(&User{})
}
