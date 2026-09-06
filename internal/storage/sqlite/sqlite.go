package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sajinlama/go-backend/internal/config"
	"github.com/sajinlama/go-backend/internal/types"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 name TEXT ,
	 email TEXT ,
	 age INTEGER
	)`)

	if err != nil {
		return nil, err
	}
	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudents(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO students (name , email , age ) VALUES ($1,$2,$3)")

	if err != nil {
		return 0, nil
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}
	lastid, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return lastid, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Students, error) {
	stmt, err := s.Db.Prepare("  SELECT id, name, email, age FROM students WHERE id = $1")
	if err != nil {
		return types.Students{}, err
	}
	defer stmt.Close()

	var students types.Students
	err = stmt.QueryRow(id).Scan(&students.Id, &students.Name, &students.Email, &students.Age)
	if err != nil {

		if err == sql.ErrNoRows {
			return types.Students{}, fmt.Errorf("no students fround with id %v", fmt.Sprint(id))
		}
		return types.Students{}, fmt.Errorf("query not found :%w", err)
	}
	return students, nil
}
