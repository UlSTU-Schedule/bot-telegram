package postgres

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// TODO: var _ store.StudentStore = (*StudentStore)(nil)

type StudentStore struct {
	db                *sqlx.DB
	studentRepository *StudentRepository
}

func NewStudentStore(db *sqlx.DB) *StudentStore {
	return &StudentStore{
		db: db,
	}
}

func (s *StudentStore) Student() *StudentRepository {
	if s.studentRepository != nil {
		return s.studentRepository
	}

	s.studentRepository = &StudentRepository{
		store: s,
	}

	return s.studentRepository
}
