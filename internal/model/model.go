package model

// Student ...
type Student struct {
	ID        int
	UserID    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	GroupName string `db:"group_name"`
	FacultyID byte   `db:"faculty_id"`
}
