package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/meshyampratap01/fileUploader/internal/services"
)

type FileHandler struct {
	fileService *services.FileService
}

func NewFileHandler(filesvc *services.FileService) *FileHandler {
	return &FileHandler{
		fileService: filesvc,
	}
}

func (fh *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := fh.fileService.SaveFiletoDB(file, header.Filename); err != nil {
		newErr := fmt.Sprintf("unable to save file: %v", err)
		http.Error(w, newErr, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File uploaded successfully")
}

func (fh *FileHandler) DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.PathValue("id")

	fileID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid file ID", http.StatusBadRequest)
		return
	}

	file, err := fh.fileService.GetFileByID(fileID)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\""+file.Name+"\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(file.Data)
}
