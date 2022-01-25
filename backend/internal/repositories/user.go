package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (repository *UserRepository) FindAll() ([]*models.User, error) {
	var users []*models.User
	err := repository.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *UserRepository) Find(id uint) (*models.User, error) {
	var users *models.User
	err := repository.db.First(&users, id).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *UserRepository) Create(user *models.User) (*models.User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) Update(user *models.User) (*models.User, error) {
	err := repository.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) Delete(id uint) (*models.User, error) {
	var user, err = repository.Find(id)
	if err != nil {
		return nil, err
	}
	err = repository.db.Delete(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
