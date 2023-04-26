package repository

import (
	"main/models"

	"gorm.io/gorm"
)

type ShoesRepository interface {
	CreateShoes(data models.Shoes) error
	GetAllShoes() ([]models.Shoes, error)
	GetDetailShoes(id int) (models.Shoes, error)
	UpdateShoes(id int, data models.Shoes) error
	DeleteShoes(id int) error
}
type shoesRepository struct {
	db *gorm.DB
}

func NewShoesRepository(db *gorm.DB) *shoesRepository {
	return &shoesRepository{db}
}

func (r *shoesRepository) CreateShoes(data models.Shoes) error {
	err := r.db.Create(&data).Error
	return err
}

func (r *shoesRepository) GetAllShoes() ([]models.Shoes, error) {
	shoes := make([]models.Shoes, 0)

	if err := r.db.Limit(10).Find(&shoes).Error; err != nil {
		return shoes, err

	}
	return shoes, nil
}

func (r *shoesRepository) GetDetailShoes(id int) (models.Shoes, error) {
	var shoes models.Shoes

	if err := r.db.Preload("ShoesDetail").Find(&shoes, id).Error; err != nil {
		return shoes, err
	}
	return shoes, nil
}

func (r *shoesRepository) UpdateShoes(id int, data models.Shoes) error {
	var shoesDetail = data.ShoesDetail

	if err := r.db.Model(&models.Shoes{}).Where("ID = ?", id).Updates(data).Error; err != nil {
		return err
	}

	if err := r.db.Model(&models.ShoesDetail{}).Where("shoes_id = ?", id).Updates(shoesDetail).Error; err != nil {
		return err

	}

	return nil
}

func (r *shoesRepository) DeleteShoes(id int) error {
	var shoes models.Shoes
	if err := r.db.Delete(&shoes, id).Error; err != nil {
		return err
	}
	return nil
}
