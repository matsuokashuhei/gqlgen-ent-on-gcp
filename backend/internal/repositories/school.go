package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type SchoolRepository struct {
	db *gorm.DB
}

func (r *SchoolRepository) FindAll() ([]*models.School, error) {
	var schools []*models.School
	err := r.db.Find(&schools).Error
	if err != nil {
		return nil, err
	}
	return schools, nil
}

func (r *SchoolRepository) Find(id uint) (*models.School, error) {
	var school *models.School
	err := r.db.First(&school, id).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (r *SchoolRepository) Create(school *models.School) (*models.School, error) {
	err := r.db.Create(&school).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (r *SchoolRepository) Update(school *models.School) (*models.School, error) {
	err := r.db.Save(school).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (r *SchoolRepository) Delete(id uint) (*models.School, error) {
	var school, err = r.Find(id)
	if err != nil {
		return nil, err
	}
	err = r.db.Delete(&school, id).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func NewSchoolRepository(db *gorm.DB) *SchoolRepository {
	return &SchoolRepository{db: db}
}
