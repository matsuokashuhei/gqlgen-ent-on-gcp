package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type StudioRepository struct {
	db *gorm.DB
}

func (r *StudioRepository) FindAll(schoolID uint) ([]*models.Studio, error) {
	var studios []*models.Studio
	err := r.db.Where("school_id = ?", schoolID).Find(&studios).Error
	if err != nil {
		return nil, err
	}
	return studios, nil
}

// func (r *StudioRepository) FindAll(school *models.School) ([]*models.Studio, error) {
// 	var studios []*models.Studio
// 	err := r.db.Model(&school).Association("Studios").Find(&studios)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return studios, nil
// }

func (r *StudioRepository) Find(id uint) (*models.Studio, error) {
	var studios *models.Studio
	err := r.db.First(&studios, id).Error
	if err != nil {
		return nil, err
	}
	return studios, nil
}

func (r *StudioRepository) Create(studio *models.Studio) (*models.Studio, error) {
	err := r.db.Create(&studio).Error
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (r *StudioRepository) Update(studio *models.Studio) (*models.Studio, error) {
	err := r.db.Save(studio).Error
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (r *StudioRepository) Delete(id uint) (*models.Studio, error) {
	var studio, err = r.Find(id)
	if err != nil {
		return nil, err
	}
	err = r.db.Delete(&studio, id).Error
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func NewStudioRepository(db *gorm.DB) *StudioRepository {
	return &StudioRepository{db: db}
}
