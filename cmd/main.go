package main

import (
	"github.com/meshyampratap01/fileUploader/internal/app"
	"github.com/meshyampratap01/fileUploader/db"
)

func main() {
	db := db.InitDB()

	app:=app.NewApp(db)

	app.Run()
}