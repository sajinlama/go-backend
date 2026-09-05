package storage

type Storage interface {
	CreateStudents(name string, email string, age int) (int64, error)
}
