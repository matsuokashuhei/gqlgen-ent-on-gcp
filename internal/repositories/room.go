package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func (repository *RoomRepository) FindAll(studioID uint) ([]*models.Room, error) {
	var rooms []*models.Room
	err := repository.db.Where("studio_id = ?", studioID).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (repository *RoomRepository) Find(id uint) (*models.Room, error) {
	var rooms *models.Room
	err := repository.db.First(&rooms, id).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (repository *RoomRepository) Create(room *models.Room) (*models.Room, error) {
	err := repository.db.Create(&room).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (repository *RoomRepository) Update(room *models.Room) (*models.Room, error) {
	err := repository.db.Save(room).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (repository *RoomRepository) Delete(id uint) (*models.Room, error) {
	var room, err = repository.Find(id)
	if err != nil {
		return nil, err
	}
	err = repository.db.Delete(&room, id).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db: db}
}
