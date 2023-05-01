package controllers

import (
	"main/dto"
	"main/models"
	"main/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ShoesController interface{}

type shoesController struct {
	shoesCase usecase.ShoesUsecase
	authCase  usecase.AuthUsecase
}

func NewShoesController(shoesUsecase usecase.ShoesUsecase, authCase usecase.AuthUsecase) *shoesController {
	return &shoesController{shoesCase: shoesUsecase, authCase: authCase}
}

func (s *shoesController) CreateShoes(c echo.Context) error {
	var data dto.Shoes

	userId := s.authCase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	err := s.shoesCase.CreateShoes(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success Craeted",
	})

}

func (s *shoesController) GetAllShoes(c echo.Context) error {
	shoes, err := s.shoesCase.GetAllShoes()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.ShoesListResponse{
		Status:  http.StatusOK,
		Message: "Success get shoes",
		Data:    shoes,
	})
}

func (s *shoesController) GetDetailShoes(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	shoes, err := s.shoesCase.GetDetailShoes(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.ShoesDetailResponse{
		Status:  http.StatusOK,
		Message: "Success get detail shoes",
		Data:    shoes,
	})

}

func (s *shoesController) UpdateShoes(c echo.Context) error {
	var data dto.Shoes

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userId := s.authCase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	if err := s.shoesCase.UpdateShoes(id, data); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success Updated",
	})

}

func (s *shoesController) DeleteShoes(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userId := s.authCase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	if err := s.shoesCase.DeleteShoes(id); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success Deleted",
	})
}