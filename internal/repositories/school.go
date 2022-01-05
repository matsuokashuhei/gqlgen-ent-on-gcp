package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type SchoolRepository struct {
	database *gorm.DB
}

func (repository *SchoolRepository) FindAll() []*models.School {
	var schools []*models.School
	repository.database.Find(&schools)
	return schools
}

func (repository *SchoolRepository) Create(school *models.School) (*models.School, error) {
	err := repository.database.Create(&school).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func NewSchoolRepository(database *gorm.DB) *SchoolRepository {
	return &SchoolRepository{database: database}
}
