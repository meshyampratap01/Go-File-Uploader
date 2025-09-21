package services

import (
	"fmt"
	"io"

	"github.com/meshyampratap01/fileUploader/internal/repository"
)

type FileService struct {
	fileRepo *repository.FileRepository
}

func NewFileService(repo *repository.FileRepository) *FileService {
	return &FileService{
		fileRepo: repo,
	}
}

func (fs *FileService) SaveFiletoDB(fileContent io.Reader,filename string) (error){
	data,err:=io.ReadAll(fileContent)
	if err!=nil{
		return fmt.Errorf("unable to read file content: %v",err)
	}
	return fs.fileRepo.SaveFile(data,filename)
}

// func (fs *FileService) GetFileFromDB(id int) (