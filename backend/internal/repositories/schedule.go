package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	db *gorm.DB
}

func (r *ScheduleRepository) FindAll(roomID uint) ([]*models.Schedule, error) {
	var schedules []*models.Schedule
	if err := r.db.
		Order("day_of_week").
		Order("start_time").
		Where("room_id = ?", roomID).
		Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func NewScheduleRepository(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}
