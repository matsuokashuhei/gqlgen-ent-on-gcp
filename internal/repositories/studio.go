package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type StudioRepository struct {
	database *gorm.DB
}

func (repository *StudioRepository) FindAll(schoolID uint) ([]*models.Studio, error) {
	var studios []*models.Studio
	err := repository.database.Where("school_id = ?", schoolID).Find(&studios).Error
	if err != nil {
		return nil, err
	}
	return studios, nil
}

func (repository *StudioRepository) Find(id uint) (*models.Studio, error) {
	var studios *models.Studio
	err := repository.database.First(&studios, id).Error
	if err != nil {
		return nil, err
	}
	return studios, nil
}

func (repository *StudioRepository) Create(studios *models.Studio) (*models.Studio, error) {
	err := repository.database.Create(&studios).Error
	if err != nil {
		return nil, err
	}
	return studios, nil
}

func (repsitory *StudioRepository) Update(studios *models.Studio) (*models.Studio, error) {
	err := repsitory.database.Save(studios).Error
	if err != nil {
		return nil, err
	}
	return studios, nil
}

func (repsitory *StudioRepository) Delete(studios *models.Studio) (*models.Studio, error) {
	err := repsitory.database.Delete(studios, studios.ID).Error
	if err != nil {
		return nil, err
	}
	return studios, nil
}

func NewStudioRepository(database *gorm.DB) *StudioRepository {
	return &StudioRepository{database: database}
}
