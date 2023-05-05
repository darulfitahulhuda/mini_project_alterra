package repository

import (
	"main/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(data models.Carts) error
	GetAllCarts(userId int) ([]models.Carts, error)
	UpdateCart(id int, data models.Carts) (models.Carts, error)
	DeleteCartItem(id int) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) CreateCart(data models.Carts) error {
	err := r.db.Create(&data).Error
	return err
}

func (r *cartRepository) GetAllCarts(userId int) ([]models.Carts, error) {
	carts := make([]models.Carts, 0)

	if err := r.db.Where("user_id = ?", userId).Preload("Shoes").Find(&carts).Error; err != nil {
		return carts, err
	}
	return carts, nil

}

func (r *cartRepository) UpdateCart(id int, data models.Carts) (models.Carts, error) {
	var cart models.Carts

	if err := r.db.Model(&cart).Where("ID = ?", id).Updates(data).Error; err != nil {
		return models.Carts{}, err
	}
	return cart, nil
}
func (r *cartRepository) DeleteCartItem(id int) error {
	if err := r.db.Delete(&models.Carts{}, id).Error; err != nil {
		return err
	}
	return nil
}
