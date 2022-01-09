package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type StudioRepository struct {
	db *gorm.DB
}

func (repository *StudioRepository) FindAll(schoolID uint) ([]*models.Studio, error) {
	var studios []*models.Studio
	err := repository.db.Where("school_id = ?", schoolID).Find(&studios).Error
	if err != nil {
		return nil, err
	}
	return studios, nil
}

// func (repository *StudioRepository) FindAll(school *models.School) ([]*models.Studio, error) {
// 	var studios []*models.Studio
// 	err := repository.db.Model(&school).Association("Studios").Find(&studios)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return studios, nil
// }

func (repository *StudioRepository) Find(id uint) (*models.Studio, error) {
	var studios *models.Studio
	err := repository.db.First(&studios, id).Error
	if err != nil {
		return nil, err
	}
	return studios, nil
}

func (repository *StudioRepository) Create(studio *models.Studio) (*models.Studio, error) {
	err := repository.db.Create(&studio).Error
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (repository *StudioRepository) Update(studio *models.Studio) (*models.Studio, error) {
	err := repository.db.Save(studio).Error
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (repository *StudioRepository) Delete(id uint) (*models.Studio, error) {
	var studio, err = repository.Find(id)
	if err != nil {
		return nil, err
	}
	err = repository.db.Delete(&studio, id).Error
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func NewStudioRepository(db *gorm.DB) *StudioRepository {
	return &StudioRepository{db: db}
}
