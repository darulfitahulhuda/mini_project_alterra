package controllers

import (
	"main/dto"
	"main/models"
	"main/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	userCase usecase.UserUsecase
	authCase usecase.AuthUsecase
}

func NewUserController(userUsecase usecase.UserUsecase, authUsecase usecase.AuthUsecase) *userController {
	return &userController{userCase: userUsecase, authCase: authUsecase}
}

func (u *userController) Create(c echo.Context) error {
	var data models.User

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user, err := u.userCase.CreateUser(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	token, err := u.authCase.CreateToken(dto.CreateToken{
		Id:       int(user.ID),
		UserType: user.UserType,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})

	}

	return c.JSON(http.StatusCreated,
		models.HttpResponse{
			Status:  http.StatusCreated,
			Message: "Success Create User",
			Data: models.AuthResponse{
				ID:       int(user.ID),
				Name:     user.Email,
				Email:    user.Email,
				UserType: user.UserType,
				Token:    token,
			},
		},
	)
}

func (u *userController) LoginUser(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	users, err := u.userCase.LoginUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Email or password is wrong",
		})
	}

	token, err := u.authCase.CreateToken(dto.CreateToken{
		Id:       int(users.ID),
		UserType: user.UserType,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK,
		models.HttpResponse{
			Status:  http.StatusOK,
			Message: "Success Login User",
			Data: models.AuthResponse{
				ID:       int(users.ID),
				Token:    token,
				Name:     users.Email,
				Email:    users.Email,
				UserType: users.UserType,
			},
		},
	)
}

func (u *userController) GetUserByAuth(c echo.Context) error {
	userId := u.authCase.ExtractTokenUserId(c, models.User_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized, models.HttpResponse{
			Status:  http.StatusUnauthorized,
			Message: "Token Unauthorized",
		})
	}

	user, err := u.userCase.GetUserById(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success get user",
		Data: models.UserResponse{
			ID:          int(user.ID),
			Name:        user.Name,
			Email:       user.Email,
			DateOfBirth: user.DateOfBirth,
			UserType:    user.UserType,
			Address:     user.Address,
		},
	})

}

func (u *userController) UpdateUser(c echo.Context) error {
	var data models.User

	userId := u.authCase.ExtractTokenUserId(c, models.User_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized, models.HttpResponse{
			Status:  http.StatusUnauthorized,
			Message: "Token Unauthorized",
		})
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user, err := u.userCase.UpdateUser(userId, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success update user",
		Data: models.UserResponse{
			ID:          int(user.ID),
			Name:        user.Name,
			Email:       user.Email,
			DateOfBirth: user.DateOfBirth,
			UserType:    user.UserType,
			Address:     user.Address,
		},
	})

}
func (u *userController) DeleteUser(c echo.Context) error {
	userId := u.authCase.ExtractTokenUserId(c, models.User_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized, models.HttpResponse{
			Status:  http.StatusUnauthorized,
			Message: "Token Unauthorized",
		})
	}

	_, err := u.userCase.DeleteUser(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, models.HttpResponse{
		Status:  http.StatusBadRequest,
		Message: "Success Delete user",
	})
}
