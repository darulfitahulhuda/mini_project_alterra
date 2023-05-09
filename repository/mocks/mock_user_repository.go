package mocksRepository

import (
	"main/models"

	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) Create(data models.User) (models.User, error) {
	ret := m.Called(data)

	return ret.Get(0).(models.User), ret.Error(1)
}

func (m *mockUserRepository) GetAllUsers() ([]models.User, error) {
	ret := m.Called()

	return ret.Get(0).([]models.User), ret.Error(1)
}

func (m *mockUserRepository) LoginUser(data models.User) (models.User, error) {
	ret := m.Called(data)

	return ret.Get(0).(models.User), ret.Error(1)
}

func (m *mockUserRepository) GetUserById(id int) (models.User, error) {
	ret := m.Called(id)

	return ret.Get(0).(models.User), ret.Error(1)
}

func (m *mockUserRepository) DeleteUser(id int) (models.User, error) {
	ret := m.Called(id)

	return ret.Get(0).(models.User), ret.Error(1)
}

func (m *mockUserRepository) UpdateUser(id int, data models.User) error {
	ret := m.Called(id, data)

	return ret.Error(0)
}
