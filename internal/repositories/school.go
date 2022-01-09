package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type SchoolRepository struct {
	db *gorm.DB
}

func (repository *SchoolRepository) FindAll() ([]*models.School, error) {
	var schools []*models.School
	err := repository.db.Find(&schools).Error
	if err != nil {
		return nil, err
	}
	return schools, nil
}

func (repository *SchoolRepository) Find(id uint) (*models.School, error) {
	var school *models.School
	err := repository.db.First(&school, id).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (repository *SchoolRepository) Create(school *models.School) (*models.School, error) {
	err := repository.db.Create(&school).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (repository *SchoolRepository) Update(school *models.School) (*models.School, error) {
	err := repository.db.Save(school).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (repository *SchoolRepository) Delete(id uint) (*models.School, error) {
	var school, err = repository.Find(id)
	if err != nil {
		return nil, err
	}
	err = repository.db.Delete(&school, id).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func NewSchoolRepository(db *gorm.DB) *SchoolRepository {
	return &SchoolRepository{db: db}
}
