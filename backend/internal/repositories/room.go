package repositories

import (
	"github.com/matsuokashuhei/landin/internal/models"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func (r *RoomRepository) FindAll(studioID uint) ([]*models.Room, error) {
	var rooms []*models.Room
	err := r.db.Where("studio_id = ?", studioID).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomRepository) Find(id uint) (*models.Room, error) {
	var rooms *models.Room
	err := r.db.First(&rooms, id).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomRepository) Create(room *models.Room) (*models.Room, error) {
	err := r.db.Create(&room).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *RoomRepository) Update(room *models.Room) (*models.Room, error) {
	err := r.db.Save(room).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *RoomRepository) Delete(id uint) (*models.Room, error) {
	var room, err = r.Find(id)
	if err != nil {
		return nil, err
	}
	err = r.db.Delete(&room, id).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db: db}
}
