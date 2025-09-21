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

func (fr *FileRepository) SaveFile(fileContent []byte, filename string) (int, error) {
	stmt, err := fr.db.Prepare("INSERT INTO files (data,name) VALUES (?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(fileContent, filename)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *FileRepository) GetFileByID(id int) (*models.File, error) {
	row := r.db.QueryRow("SELECT id, data, name FROM files WHERE id = ?", id)

	var f models.File
	err := row.Scan(&f.ID, &f.Data, &f.Name)
	if err == sql.ErrNoRows {
		return nil, errors.New("file not found")
	} else if err != nil {
		return nil, errors.New("failed to scan file row: " + err.Error())
	}

	return &f, nil
}
