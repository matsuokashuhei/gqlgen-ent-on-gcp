package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type InstructorRepository struct {
	db *gorm.DB
}

func (r *InstructorRepository) FindAll() ([]*models.Instructor, error) {
	var instructors []*models.Instructor
	if err := r.db.Find(&instructors).Error; err != nil {
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

func NewInstructorRepository(db *gorm.DB) *InstructorRepository {
	return &InstructorRepository{db: db}
}
