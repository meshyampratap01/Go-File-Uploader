package repository

import (
	"database/sql"
	"errors"

	"github.com/meshyampratap01/fileUploader/internal/models"
)

type FileRepository struct {
	db *sql.DB
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

func (fr *FileRepository) SaveFile(fileContent []byte, filename string) error {
	stmt, err := fr.db.Prepare("INSERT INTO files (data,name) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(fileContent, filename)
	return err
}

func (r *FileRepository) GetFileByID(id int) (*models.File, error) {
	row := r.db.QueryRow("SELECT id, content, name FROM files WHERE id = ?", id)

	var f models.File
	err := row.Scan(&f.ID, &f.Data, &f.Name)
	if err == sql.ErrNoRows {
		return nil, errors.New("file not found")
	} else if err != nil {
		return nil, err
	}

	return &f, nil
}
