package usecase

import (
	"errors"
	"main/dto"
	"main/models"
	"main/repository"

	"golang.org/x/crypto/bcrypt"
)

// const admin = "admin"
const user = "user"

type UserUsecase interface {
	CreateUser(payload dto.CreateUser) (models.User, error)
	GetAllUsers() ([]models.User, error)
	LoginUser(payload dto.LoginUser) (models.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (s *userUsecase) CreateUser(payload dto.CreateUser) (models.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		return models.User{}, err
	}

	data := models.User{
		Name:        payload.Name,
		Email:       payload.Email,
		Password:    string(password),
		DateOfBirth: payload.DateOfBirth,
		UserType:    user,
	}
	user, err := s.userRepository.Create(data)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userUsecase) GetAllUsers() ([]models.User, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return users, err

	}
	return users, nil
}

func (s *userUsecase) LoginUser(payload dto.LoginUser) (models.User, error) {
	data := models.User{
		Email:    payload.Email,
		Password: payload.Password,
	}
	user, err := s.userRepository.LoginUser(data)

	if err != nil {
		return user, err
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if errPassword != nil {
		return user, errors.New("password not same")
	}

	return user, nil
}
