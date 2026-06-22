package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// 環境変数確認用
	fmt.Println("===== DB CONFIG =====")
	fmt.Println("DB_HOST =", host)
	fmt.Println("DB_PORT =", port)
	fmt.Println("DB_USER =", user)
	fmt.Println("DB_NAME =", dbname)
	fmt.Println("=====================")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	fmt.Println("DSN =", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("DB CONNECTION ERROR:", err)
		panic(err)
	}

	fmt.Println("✅ MySQL Connected!")

	DB = db
}
