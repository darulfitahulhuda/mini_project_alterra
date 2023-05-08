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
	var data dto.ShoesRequest

	userId := s.authCase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	shoes, err := s.shoesCase.CreateShoes(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success create shoes",
		Data:    shoes,
	})

}

func (s *shoesController) CreateShoesSize(c echo.Context) error {
	var data dto.ShoesSizeRequest

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	userId := s.authCase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	size, err := s.shoesCase.CreateShoesSize(data)

	if err != nil {

		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.HttpResponse{
		Message: "Success create shoes size",
		Status:  http.StatusCreated,
		Data:    size,
	})

}

func (s *shoesController) GetAllShoes(c echo.Context) error {
	gender := c.QueryParam("gender")

	shoes, err := s.shoesCase.GetAllShoes(gender)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success get shoes",
		Data:    shoes,
	})
}

func (s *shoesController) GetDetailShoes(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	shoes, err := s.shoesCase.GetDetailShoes(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success get detail shoes",
		Data:    shoes,
	})

}

func (s *shoesController) UpdateShoes(c echo.Context) error {
	var data dto.ShoesRequest

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userId := s.authCase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	if err := s.shoesCase.UpdateShoes(id, data); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	shoes, err := s.shoesCase.GetDetailShoes(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success update shoes",
		Data:    shoes,
	})

}

func (s *shoesController) DeleteShoes(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userId := s.authCase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	if err := s.shoesCase.DeleteShoes(id); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success Deleted",
	})
}

func (s *shoesController) DeleteShoesSize(c echo.Context) error {
	var data dto.ShoesSizeRequest

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	userId := s.authCase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	if err := s.shoesCase.DeleteShoesSize(data); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success Deleted Shoes Size",
	})
}
