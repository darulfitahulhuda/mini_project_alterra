package controllers

import (
	"main/dto"
	"main/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController interface{}

type adminController struct {
	userCase usecase.UserUsecase
	authCase usecase.AuthUsecase
}

func NewAdminController(userUsecase usecase.UserUsecase, authUsecase usecase.AuthUsecase) *adminController {
	return &adminController{userCase: userUsecase, authCase: authUsecase}
}

func (a *adminController) GetAllUsers(c echo.Context) error {
	userId := a.authCase.ExtractTokenAdminId(c)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"status":  http.StatusUnauthorized,
			"message": "Token Unauthorized",
		})
	}

	users, err := a.userCase.GetAllUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success Get All User",
		"data":    users,
	})

}

func (a *adminController) GetAdminByAuth(c echo.Context) error {
	adminId := a.authCase.ExtractTokenAdminId(c)

	if adminId == 0 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"status":  http.StatusUnauthorized,
			"message": "Token Unauthorized",
		})
	}

	admin, err := a.userCase.GetUserById(adminId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Success get admin",
		"data":    admin,
	})
}

func (a *adminController) UpdateAdmin(c echo.Context) error {
	var data dto.UpdateUser

	adminId := a.authCase.ExtractTokenAdminId(c)

	if adminId == 0 {
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

	admin, err := a.userCase.UpdateUser(adminId, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Success update admin",
		"data":    admin,
	})

}
func (a *adminController) DeleteAdmin(c echo.Context) error {
	adminId := a.authCase.ExtractTokenAdminId(c)

	if adminId == 0 {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"status":  http.StatusUnauthorized,
			"message": "Token Unauthorized",
		})
	}

	_, err := a.userCase.DeleteUser(adminId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Success Delete admin",
	})

}
