package usecase

import (
	"errors"
	"main/models"
	mocksRepository "main/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// go test ./usecase -cover
// go test ./usecase/ -coverprofile=cover.out && go tool cover -html=cover.out

func TestLoaginSuccess(t *testing.T) {
	mockUser := models.User{
		Email:    "user1@gmail.com",
		Password: "12345678",
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(mockUser.Password), 14)

	returnUser := models.User{
		Email:    "user1@gmail.com",
		Password: string(password),
	}

	mockUserRepository := mocksRepository.NewMockUserRepository()
	mockUserRepository.On("LoginUser", mockUser).Return(returnUser, nil)
	service := NewUserUsecase(mockUserRepository)

	user, err := service.LoginUser(mockUser)

	if err != nil {
		t.Errorf("Got Error %v", err)
	}

	assert.Equal(t, user.Email, mockUser.Email)
}

func TestLoaginFailedPasswordNotSame(t *testing.T) {
	mockUser := models.User{
		Email:    "user1@gmail.com",
		Password: "12345678",
	}

	mockUserRepository := mocksRepository.NewMockUserRepository()
	mockUserRepository.On("LoginUser", mockUser).Return(mockUser, errors.New("failed"))
	service := NewUserUsecase(mockUserRepository)

	user, err := service.LoginUser(mockUser)

	if err != nil {
		assert.Error(t, err, "failed")
	}

	assert.Equal(t, user.Email, mockUser.Email)
}

func TestUserGetAllUsersSuccess(t *testing.T) {
	mockUser := make([]models.User, 0)

	mockUser = append(mockUser, models.User{
		Email: "test1@gmail.com",
	})

	mockUserRepository := mocksRepository.NewMockUserRepository()

	mockUserRepository.On("GetAllUsers").Return(mockUser, nil)

	service := NewUserUsecase(mockUserRepository)

	users, err := service.GetAllUsers()

	if err != nil {
		t.Errorf("Got Error %v", err)
	}

	assert.Equal(t, users[0].Email, mockUser[0].Email)
}

func TestUserGetAllUsersFailed(t *testing.T) {

	mockUser := make([]models.User, 0)

	mockUser = append(mockUser, models.User{
		Email: "test1@gmail.com",
	})

	mockUserRepository := mocksRepository.NewMockUserRepository()
	mockUserRepository.On("GetAllUsers").Return(mockUser, errors.New("failed"))

	service := NewUserUsecase(mockUserRepository)

	users, err := service.GetAllUsers()

	if err != nil {
		// t.Errorf("Got Error %v", err)
		assert.Error(t, err, "failed")
	}

	assert.Equal(t, users[0].Email, mockUser[0].Email)
}

func TestUserGetByIDSuccess(t *testing.T) {
	mockUser := models.User{
		Email: "test1@gmail.com",
	}

	id := 1

	mockUserRepository := mocksRepository.NewMockUserRepository()

	mockUserRepository.On("GetUserById", id).Return(mockUser, nil)

	service := NewUserUsecase(mockUserRepository)

	users, err := service.GetUserById(id)

	if err != nil {
		t.Errorf("Got Error %v", err)
	}

	assert.Equal(t, users.Email, mockUser.Email)
}

func TestUserGetByIDUsersFailed(t *testing.T) {

	mockUser := models.User{
		Email: "test1@gmail.com",
	}

	id := 1

	mockUserRepository := mocksRepository.NewMockUserRepository()
	mockUserRepository.On("GetUserById", id).Return(mockUser, errors.New("failed"))

	service := NewUserUsecase(mockUserRepository)

	_, err := service.GetUserById(id)

	if err != nil {
		// t.Errorf("Got Error %v", err)
		assert.Error(t, err, "failed")
	}

}

func TestUpdateUserSuccess(t *testing.T) {
	mockUser := models.User{
		Email: "test1@gmail.com",
	}

	id := 1

	mockUserRepository := mocksRepository.NewMockUserRepository()

	mockUserRepository.On("UpdateUser", id, mockUser).Return(nil)
	mockUserRepository.On("GetUserById", id).Return(mockUser, nil)

	service := NewUserUsecase(mockUserRepository)

	users, err := service.UpdateUser(id, mockUser)

	if err != nil {
		t.Errorf("Got Error %v", err)
	}

	assert.Equal(t, users.Email, mockUser.Email)
}

func TestUpdateUserFailed(t *testing.T) {

	mockUser := models.User{
		Email: "test1@gmail.com",
	}

	id := 1

	mockUserRepository := mocksRepository.NewMockUserRepository()
	mockUserRepository.On("UpdateUser", id, mockUser).Return(errors.New("failed"))

	service := NewUserUsecase(mockUserRepository)

	_, err := service.UpdateUser(id, mockUser)

	if err != nil {
		// t.Errorf("Got Error %v", err)
		assert.Error(t, err, "failed")
	}

}

func TestUpdateUserIdFailed(t *testing.T) {

	mockUser := models.User{
		Email: "test1@gmail.com",
	}

	id := 1

	mockUserRepository := mocksRepository.NewMockUserRepository()
	mockUserRepository.On("UpdateUser", id, mockUser).Return(nil)
	mockUserRepository.On("GetUserById", id).Return(mockUser, errors.New("failed"))

	service := NewUserUsecase(mockUserRepository)

	_, err := service.UpdateUser(id, mockUser)

	if err != nil {
		// t.Errorf("Got Error %v", err)
		assert.Error(t, err, "failed")
	}

}

func TestDeletUserSuccess(t *testing.T) {

	id := 1

	mockUserRepository := mocksRepository.NewMockUserRepository()

	mockUserRepository.On("DeleteUser", id).Return(models.User{}, nil)

	service := NewUserUsecase(mockUserRepository)

	_, err := service.DeleteUser(id)

	if err != nil {
		t.Errorf("Got Error %v", err)
	}

}

func TestDeletUserFailed(t *testing.T) {

	id := 1

	mockUserRepository := mocksRepository.NewMockUserRepository()
	mockUserRepository.On("DeleteUser", id).Return(models.User{}, errors.New("failed"))

	service := NewUserUsecase(mockUserRepository)

	_, err := service.DeleteUser(id)

	if err != nil {
		assert.Error(t, err, "failed")
	}

}
