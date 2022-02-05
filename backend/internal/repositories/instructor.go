package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type InstructorRepository struct {
	db *gorm.DB
}

func (r *InstructorRepository) FindAll(offset int, limit int) ([]*models.Instructor, error) {
	var instructors []*models.Instructor
	if err := r.db.Offset(offset).Limit(limit).Find(&instructors).Error; err != nil {
		return nil, err
	}
	return instructors, nil
}

func (r *InstructorRepository) Find(id uint) (*models.Instructor, error) {
	var instructor *models.Instructor
	if err := r.db.First(&instructor, id).Error; err != nil {
		return nil, err
	}
	return instructor, nil
}

func (r *InstructorRepository) CountAll() (*int64, error) {
	var count *int64
	if err := r.db.Model(&models.Instructor{}).Count(count).Error; err != nil {
		return nil, err
	}
	return count, nil
}

func NewInstructorRepository(db *gorm.DB) *InstructorRepository {
	return &InstructorRepository{db: db}
}
