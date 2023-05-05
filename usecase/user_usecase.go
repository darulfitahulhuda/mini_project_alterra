package usecase

import (
	"errors"
	"main/dto"
	"main/models"
	"main/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(payload dto.CreateUser) (models.User, error)
	GetAllUsers() ([]models.User, error)
	LoginUser(payload dto.LoginUser) (models.User, error)
	GetUserById(id int) (models.User, error)
	DeleteUser(id int) (models.User, error)
	UpdateUser(id int, payload dto.UpdateUser) (models.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (u *userUsecase) CreateUser(payload dto.CreateUser) (models.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		return models.User{}, err
	}

	data := models.User{
		Name:        payload.Name,
		Email:       payload.Email,
		Password:    string(password),
		DateOfBirth: payload.DateOfBirth,
		UserType:    payload.UserType,
	}
	user, err := u.userRepository.Create(data)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *userUsecase) LoginUser(payload dto.LoginUser) (models.User, error) {
	data := models.User{
		Email:    payload.Email,
		Password: payload.Password,
		UserType: payload.UserType,
	}
	user, err := u.userRepository.LoginUser(data)

	if err != nil {
		return user, err
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if errPassword != nil {
		return models.User{}, errors.New("password not same")
	}

	return user, nil
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	users, err := u.userRepository.GetAllUsers()
	if err != nil {
		return users, err

	}
	return users, nil
}

func (u *userUsecase) GetUserById(id int) (models.User, error) {
	user, err := u.userRepository.GetUserById(id)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *userUsecase) UpdateUser(id int, payload dto.UpdateUser) (models.User, error) {
	user := models.User{
		Name:        payload.Name,
		Email:       payload.Email,
		DateOfBirth: payload.DateOfBirth,
		Address:     payload.Address,
	}

	err := u.userRepository.UpdateUser(id, user)

	if err != nil {
		return models.User{}, err
	}

	getUser, err := u.userRepository.GetUserById(id)

	if err != nil {
		return models.User{}, err
	}

	return getUser, nil
}

func (u *userUsecase) DeleteUser(id int) (models.User, error) {
	user, err := u.userRepository.DeleteUser(id)

	if err != nil {
		return models.User{}, err

	}

	return user, nil
}
