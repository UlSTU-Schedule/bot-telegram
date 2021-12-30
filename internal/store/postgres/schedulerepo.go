package postgres

import (
	"database/sql"
	"fmt"
	"github.com/ulstu-schedule/bot-telegram/internal/model"
	"github.com/ulstu-schedule/bot-telegram/internal/store"
)

const groupScheduleRepoName = "groups_schedule"

var _ store.GroupScheduleRepository = (*GroupScheduleRepository)(nil)

type GroupScheduleRepository struct {
	store *ScheduleStore
}

func (r *GroupScheduleRepository) GetSchedule(groupName string) (*model.GroupSchedule, error) {
	var schedule model.GroupSchedule
	query := fmt.Sprintf("SELECT * FROM %s WHERE group_name=$1", groupScheduleRepoName)
	err := r.store.db.Get(&schedule, query, groupName)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// if the group schedule is not in the database
	if schedule.Name == "" {
		return nil, nil
	}

	return &schedule, nil
}

func (r *GroupScheduleRepository) GetGroups() ([]string, error) {
	var groups []string
	query := fmt.Sprintf("SELECT group_name FROM %s", groupScheduleRepoName)
	err := r.store.db.Select(&groups, query)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

var _ store.TeacherScheduleRepository = (*TeacherScheduleRepository)(nil)

type TeacherScheduleRepository struct {
	// TODO: сделать по примеру GroupScheduleRepository
}
