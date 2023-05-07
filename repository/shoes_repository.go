package repository

import (
	"main/models"

	"gorm.io/gorm"
)

type ShoesRepository interface {
	CreateShoes(data models.Shoes) (models.Shoes, error)
	CreateShoesSize(data models.ShoesSize) (models.ShoesSize, error)

	GetAllShoes(gender string) ([]models.Shoes, error)
	GetDetailShoes(id int) (models.Shoes, error)
	GetShoesSize(data models.ShoesSize) (models.ShoesSize, error)

	UpdateShoes(id int, data models.Shoes) error
	UpdateShoesSize(data models.ShoesSize) error
	ReduceShoesQty(data models.ShoesSize) error

	DeleteShoes(id int) error
	DeleteShoesSize(data models.ShoesSize, all bool) error
}
type shoesRepository struct {
	db *gorm.DB
}

func NewShoesRepository(db *gorm.DB) *shoesRepository {
	return &shoesRepository{db}
}

func (r *shoesRepository) CreateShoes(data models.Shoes) (models.Shoes, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *shoesRepository) CreateShoesSize(data models.ShoesSize) (models.ShoesSize, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *shoesRepository) GetAllShoes(gender string) ([]models.Shoes, error) {
	shoes := make([]models.Shoes, 0)
	if gender == "" {
		if err := r.db.Order("ID desc").Limit(10).Find(&shoes).Error; err != nil {
			return shoes, err
		}
	} else {
		if err := r.db.Where("gender = ?", gender).Order("ID desc").Limit(10).Find(&shoes).Error; err != nil {
			return shoes, err
		}

	}

	return shoes, nil
}

func (r *shoesRepository) GetDetailShoes(id int) (models.Shoes, error) {
	var shoes models.Shoes

	if err := r.db.Preload("ShoesDetail").Preload("Sizes").Find(&shoes, id).Error; err != nil {
		return models.Shoes{}, err
	}
	return shoes, nil
}
func (r *shoesRepository) GetShoesSize(data models.ShoesSize) (models.ShoesSize, error) {
	var shoesSize models.ShoesSize

	if err := r.db.Where("shoes_id = ? AND size = ?", data.ShoesId, data.Size).First(&shoesSize).Error; err != nil {
		return shoesSize, err

	}
	return shoesSize, nil
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

func (r *shoesRepository) UpdateShoesSize(data models.ShoesSize) error {
	if err := r.db.Model(&models.ShoesSize{}).Where("shoes_id = ? AND size = ?", data.ShoesId, data.Size).Update("qty", data.Qty).Error; err != nil {
		return err
	}

	return nil
}

func (r *shoesRepository) ReduceShoesQty(data models.ShoesSize) error {
	var shoesSize models.ShoesSize

	if err := r.db.Where("shoes_id = ? AND size = ?", data.ShoesId, data.Size).First(&shoesSize).Error; err != nil {
		return err
	}

	shoesSize.Qty -= data.Qty

	if err := r.db.Save(&shoesSize).Error; err != nil {
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

func (r *shoesRepository) DeleteShoesSize(data models.ShoesSize, all bool) error {
	if all {
		if err := r.db.Where("shoes_id = ?", data.ShoesId).Delete(&data).Error; err != nil {
			return err
		}
	} else {
		if err := r.db.Where("shoes_id = ? AND size = ?", data.ShoesId, data.Size).Delete(&data).Error; err != nil {
			return err
		}
	}

	return nil
}
