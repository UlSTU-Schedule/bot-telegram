package postgres

import (
	"github.com/jmoiron/sqlx"
)

// TODO: var _ store.ScheduleStore = (*ScheduleStore)(nil)

type ScheduleStore struct {
	db                        *sqlx.DB
	groupScheduleRepository   *GroupScheduleRepository
	teacherScheduleRepository *TeacherScheduleRepository
}

func NewScheduleStore(db *sqlx.DB) *ScheduleStore {
	return &ScheduleStore{
		db: db,
	}
}

func (s *ScheduleStore) GroupSchedule() *GroupScheduleRepository {
	if s.groupScheduleRepository != nil {
		return s.groupScheduleRepository
	}

	s.groupScheduleRepository = &GroupScheduleRepository{
		store: s,
	}

	return s.groupScheduleRepository
}

func (s *ScheduleStore) TeacherSchedule() *TeacherScheduleRepository {
	// TODO: сделать по примеру GroupSchedule()
	return nil
}
