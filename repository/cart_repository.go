package repository

import (
	"main/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(data models.Carts) (models.Carts, error)
	GetAllCarts(userId int) ([]models.Carts, error)
	GetCartById(id int) (models.Carts, error)
	UpdateCart(id int, data models.Carts) (models.Carts, error)
	DeleteCartItem(id int) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) CreateCart(data models.Carts) (models.Carts, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *cartRepository) GetCartById(id int) (models.Carts, error) {
	var cart models.Carts

	if err := r.db.First(&cart, id).Error; err != nil {
		return cart, err
	}
	return cart, nil

}

func (r *cartRepository) GetAllCarts(userId int) ([]models.Carts, error) {
	carts := make([]models.Carts, 0)

	if err := r.db.Where("user_id = ?", userId).Order("ID desc").Preload("Shoes").Find(&carts).Error; err != nil {
		return carts, err
	}
	return carts, nil

}

func (r *cartRepository) UpdateCart(id int, data models.Carts) (models.Carts, error) {
	var cart models.Carts

	if err := r.db.First(&cart, id).Error; err != nil {
		return models.Carts{}, err
	}
	cart.Qty = data.Qty
	cart.Size = data.Size
	cart.Status = data.Status

	if err := r.db.Save(&cart).Error; err != nil {
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
