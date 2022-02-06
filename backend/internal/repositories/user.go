package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) FindAll() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Find(id uint) (*models.User, error) {
	var user *models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByAuth(UID string) (*models.User, error) {
	var user *models.User
	err := r.db.Where("firebase_uid = ?", UID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(user *models.User) (*models.User, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Delete(id uint) (*models.User, error) {
	var user, err = r.Find(id)
	if err != nil {
		return nil, err
	}
	err = r.db.Delete(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
