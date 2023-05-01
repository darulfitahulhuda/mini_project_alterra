package repository

import (
	"main/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(data models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	LoginUser(data models.User) (models.User, error)
	GetUserById(id int) (models.User, error)
	DeleteUser(id int) (models.User, error)
	UpdateUser(id int, data models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(data models.User) (models.User, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	users := make([]models.User, 0)
	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepository) LoginUser(data models.User) (models.User, error) {
	var user models.User
	var e error
	if e = r.db.Model(&user).Where("email = ? AND user_type = ?", data.Email, data.UserType).First(&user).Error; e != nil {
		return user, e
	}
	return user, nil
}

func (r *userRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	if e := r.db.First(&user, id).Error; e != nil {
		return user, e
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id int) (models.User, error) {
	var user models.User
	if e := r.db.Delete(&user, id).Error; e != nil {
		return user, e
	}
	return user, nil

}

func (r *userRepository) UpdateUser(id int, data models.User) (models.User, error) {
	if e := r.db.Model(&data).Where("ID = ?", id).Updates(data).Error; e != nil {
		return data, e
	}
	return data, nil
}