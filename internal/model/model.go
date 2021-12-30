package model

import (
	"github.com/jmoiron/sqlx/types"
	"time"
)

// Student ...
type Student struct {
	ID        int
	UserID    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	GroupName string `db:"group_name"`
	FacultyID byte   `db:"faculty_id"`
}

// GroupSchedule ...
type GroupSchedule struct {
	ID         int
	Name       string         `db:"group_name"`
	UpdateTime time.Time      `db:"update_time"`
	Info       types.JSONText `db:"info"`
}

// TeacherSchedule ...
type TeacherSchedule struct {
	// TODO: сделать по аналогии с GroupSchedule
}
