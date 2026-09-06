package storage

import "github.com/sajinlama/go-backend/internal/types"

type Storage interface {
	CreateStudents(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Students, error)
}
