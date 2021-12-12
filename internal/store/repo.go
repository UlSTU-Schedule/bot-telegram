package store

import "github.com/ulstu-schedule/bot-telegram/internal/model"

// StudentRepository ...
type StudentRepository interface {
	// GetAllStudents ...
	GetAllStudents() ([]model.Student, error)

	// Information ...
	Information(firstName, lastName string, userID int, groupName string, facultyID byte) error

	// AddStudent ...
	AddStudent(firstName, lastName string, userID int, groupName string, facultyID byte)

	// GetStudent ...
	GetStudent(userID int) (*model.Student, error)

	// UpdateStudent ...
	UpdateStudent(firstName, lastName string, userID int, newGroupName string, facultyID byte)
}
