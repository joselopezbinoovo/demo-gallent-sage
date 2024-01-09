package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", username, password, dbHost, dbName)
	log.Printf("connected to mssql: %s/%s \n", dbHost, dbName)

	conn, err := gorm.Open("mssql", dbUri)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	db = conn
}

func GetDB() *gorm.DB {
	return db
}
