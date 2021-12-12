package postgres

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ulstu-schedule/bot-telegram/internal/model"
	"github.com/ulstu-schedule/bot-telegram/internal/store"
)

const studentsRepositoryName = "telegram_students"

var _ store.StudentRepository = (*StudentRepository)(nil)

// StudentRepository ...
type StudentRepository struct {
	DB *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

func (r *StudentRepository) GetAllStudents() ([]model.Student, error) {
	var students []model.Student
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id", studentsRepositoryName)
	err := r.DB.Select(&students, query)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (r *StudentRepository) Information(firstName, lastName string, userID int, groupName string, facultyID byte) error {
	student, err := r.GetStudent(userID)
	if err != nil {
		return err
	}

	if student != nil {
		r.UpdateStudent(firstName, lastName, userID, groupName, facultyID)
	} else {
		r.AddStudent(firstName, lastName, userID, groupName, facultyID)
	}
	return nil
}

func (r *StudentRepository) AddStudent(firstName, lastName string, userID int, groupName string, facultyID byte) {
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, user_id, group_name, faculty_id) VALUES ($1, $2, $3, $4, $5)", studentsRepositoryName)
	r.DB.MustExec(query, firstName, lastName, userID, groupName, facultyID)
}

func (r *StudentRepository) GetStudent(userID int) (*model.Student, error) {
	var student model.Student
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", studentsRepositoryName)
	err := r.DB.Get(&student, query, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// студента нет в БД (если бы студент был в бд, то в поле ID был бы его ID VK)
	if student.ID == 0 {
		return nil, nil
	}

	return &student, nil
}

func (r *StudentRepository) UpdateStudent(firstName, lastName string, userID int, newGroupName string, newFacultyID byte) {
	query := fmt.Sprintf("UPDATE %s SET first_name=$2, last_name=$3, group_name=$4, faculty_id=$5 WHERE user_id=$1", studentsRepositoryName)
	r.DB.MustExec(query, userID, firstName, lastName, newGroupName, newFacultyID)
}
