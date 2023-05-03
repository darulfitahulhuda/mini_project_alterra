package controllers

import (
	"main/dto"
	"main/models"
	"main/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartController interface{}

type cartController struct {
	cartUsecase usecase.CartUsecase
	authUsecase usecase.AuthUsecase
}

func NewCartContoller(cartUsecase usecase.CartUsecase, authUsecase usecase.AuthUsecase) *cartController {
	return &cartController{cartUsecase: cartUsecase, authUsecase: authUsecase}
}

func (cc *cartController) CreateCart(c echo.Context) error {
	var payload dto.Cart

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userId := cc.authUsecase.ExtractTokenUserId(c, models.User_Type)
	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}
	if err := cc.cartUsecase.CreateCart(userId, payload); err != nil {
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

func (cc *cartController) GetAllCarts(c echo.Context) error {
	userId := cc.authUsecase.ExtractTokenUserId(c, models.User_Type)
	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	carts, err := cc.cartUsecase.GetAllCarts(userId)
	if err != nil {

		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.CartListResponse{
		Message: "Success get all carts",
		Status:  http.StatusOK,
		Data:    carts,
	})
}

func (cc *cartController) UpdateCart(c echo.Context) error {
	var payload dto.Cart

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userId := cc.authUsecase.ExtractTokenUserId(c, models.User_Type)
	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	cart, err := cc.cartUsecase.UpdateCart(id, userId, payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.CartResponse{
		Message: "Success update cart",
		Status:  http.StatusOK,
		Data:    cart,
	})

}

func (cc *cartController) DeleteCartItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userId := cc.authUsecase.ExtractTokenUserId(c, models.User_Type)
	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	if err := cc.cartUsecase.DeleteCartItem(id); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success deleted",
	})
}

// (id int) error
