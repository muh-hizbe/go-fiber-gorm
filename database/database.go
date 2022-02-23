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
	//const POSTGRESQL = "postgresql://postgres:@localhost:5432/go_fiber_gorm?sslmode=disable&TimeZone=Asia/Jakarta"
	const POSTGRESQL = "postgres://dltummspuwbpxp:f126da1dedd437c3dd5810a26906a26a87c6707255adec077a852f385259d767@ec2-34-194-14-176.compute-1.amazonaws.com:5432/d85vsfdt0o0quq"
	dsn := POSTGRESQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Connected to database")
}
