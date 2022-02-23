package database

import (
	"fmt"
	//"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	//	USING MYSQL
	//const MYSQL = "root:@tcp(127.0.0.1:3306)/go_fiber_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := MYSQL
	//DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//	USING POSTGRESQL
	const POSTGRESQL = "postgresql://postgres:@localhost:5432/go_fiber_gorm?sslmode=disable&TimeZone=Asia/Jakarta"
	dsn := POSTGRESQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Connected to database")
}
