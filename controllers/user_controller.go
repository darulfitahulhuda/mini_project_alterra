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
		UserType: user.UserType,
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

func (u *userController) LoginUser(c echo.Context) error {
	user := dto.LoginUser{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	users, err := u.userCase.LoginUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "password or email is wrong",
		})
	}

	token, err := u.authCase.CreateToken(dto.CreateToken{
		Id:       int(users.ID),
		UserType: user.UserType,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Success Create User",
		"data":    users,
		"token":   token,
	})
}

func (u *userController) GetUserByAuth(c echo.Context) error {
	userId := u.authCase.ExtractTokenUserId(c)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"status":  http.StatusUnauthorized,
			"message": "Token Unauthorized",
		})
	}

	user, err := u.userCase.GetUserById(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Success get user",
		"data":    user,
	})
}

func (u *userController) UpdateUser(c echo.Context) error {
	var data dto.UpdateUser

	userId := u.authCase.ExtractTokenUserId(c)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"status":  http.StatusUnauthorized,
			"message": "Token Unauthorized",
		})
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	user, err := u.userCase.UpdateUser(userId, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Success update user",
		"data":    user,
	})

}
func (u *userController) DeleteUser(c echo.Context) error {
	userId := u.authCase.ExtractTokenUserId(c)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"status":  http.StatusUnauthorized,
			"message": "Token Unauthorized",
		})
	}

	_, err := u.userCase.DeleteUser(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Success Delete user",
	})

}
