package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type SchoolRepository struct {
	database *gorm.DB
}

func (repository *SchoolRepository) FindAll() ([]*models.School, error) {
	var schools []*models.School
	err := repository.database.Find(&schools).Error
	if err != nil {
		return nil, err
	}
	return schools, nil
}

func (repository *SchoolRepository) Find(id int) (*models.School, error) {
	var school *models.School
	err := repository.database.First(&school, id).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (repository *SchoolRepository) Create(school *models.School) (*models.School, error) {
	err := repository.database.Create(&school).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (repsitory *SchoolRepository) Update(school *models.School) (*models.School, error) {
	err := repsitory.database.Save(school).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (repsitory *SchoolRepository) Delete(school *models.School) (*models.School, error) {
	err := repsitory.database.Delete(school, school.ID).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func NewSchoolRepository(database *gorm.DB) *SchoolRepository {
	return &SchoolRepository{database: database}
}
