package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/meshyampratap01/fileUploader/internal/handlers"
	"github.com/meshyampratap01/fileUploader/internal/repository"
	"github.com/meshyampratap01/fileUploader/internal/services"
)

var baseURL = "/api/v1"

type App struct {
	db     *sql.DB
	apimux *http.ServeMux

	fileHandler *handlers.FileHandler
}

func NewApp(db *sql.DB) *App {
	fileRepo := repository.NewFileRepository(db)

	fileSvc := services.NewFileService(fileRepo)

	fileHandler := handlers.NewFileHandler(fileSvc)

	app := &App{
		db:          db,
		apimux:      http.NewServeMux(),
		fileHandler: fileHandler,
	}

	app.apimux.HandleFunc("POST "+baseURL+"/upload", app.fileHandler.UploadFile)
	app.apimux.HandleFunc("GET "+baseURL+"/files/{id}/download", app.fileHandler.DownloadFileHandler)

	return app
}

func (app *App) Run() {
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", app.apimux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
