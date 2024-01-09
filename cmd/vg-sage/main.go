package main

import (
	"vg-sage/internal/server"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
	server.SetupServer()
}
