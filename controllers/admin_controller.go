package controllers

import (
	"main/models"
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
	userId := a.authCase.ExtractTokenUserId(c, models.Admin_Type)

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

	userResponse := make([]models.UserResponse, 0)

	for _, v := range users {
		userResponse = append(userResponse, models.UserResponse{
			ID:          int(v.ID),
			Name:        v.Email,
			Email:       v.Email,
			DateOfBirth: v.DateOfBirth,
			Address:     v.Address,
			UserType:    v.UserType,
		})

	}

	return c.JSON(http.StatusOK,
		models.HttpResponse{
			Status:  http.StatusOK,
			Message: "Success Get All User",
			Data:    userResponse,
		},
	)

}

func (a *adminController) GetAdminByAuth(c echo.Context) error {
	adminId := a.authCase.ExtractTokenUserId(c, models.Admin_Type)

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

	return c.JSON(http.StatusOK,
		models.HttpResponse{
			Status:  http.StatusOK,
			Message: "Success Get Admin",
			Data: models.UserResponse{
				ID:          int(admin.ID),
				Name:        admin.Email,
				Email:       admin.Email,
				DateOfBirth: admin.DateOfBirth,
				Address:     admin.Address,
				UserType:    admin.UserType,
			},
		},
	)
}

func (a *adminController) UpdateAdmin(c echo.Context) error {
	var data models.User

	adminId := a.authCase.ExtractTokenUserId(c, models.Admin_Type)

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
	return c.JSON(http.StatusOK,
		models.HttpResponse{
			Status:  http.StatusOK,
			Message: "Success update admin",
			Data: models.UserResponse{
				ID:          int(admin.ID),
				Name:        admin.Email,
				Email:       admin.Email,
				DateOfBirth: admin.DateOfBirth,
				Address:     admin.Address,
				UserType:    admin.UserType,
			},
		},
	)

}
func (a *adminController) DeleteAdmin(c echo.Context) error {
	adminId := a.authCase.ExtractTokenUserId(c, models.Admin_Type)

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
