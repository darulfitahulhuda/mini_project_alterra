package controllers

import (
	"main/dto"
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
	var data dto.CreateUser

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := u.userCase.CreateUser(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := u.authCase.CreateToken(dto.CreateToken{
		Id:       int(user.ID),
		UserType: "user",
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"data":    user,
		"message": "Success Create User",
		"token":   token,
	})
}

func (u *userController) GetAllUsers(c echo.Context) error {
	users, err := u.userCase.GetAllUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success Get All User",
		"data":    users,
	})

}

func (u *userController) LoginUser(c echo.Context) error {
	user := dto.LoginUser{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	users, err := u.userCase.LoginUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "password or email is wrong",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Success Create User",
		"data":    users,
	})
}
