package controllers

import (
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
	var payload models.Carts

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userId := cc.authUsecase.ExtractTokenUserId(c, models.User_Type)
	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}
	cart, err := cc.cartUsecase.CreateCart(userId, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated,
		models.HttpResponse{
			Status:  http.StatusCreated,
			Message: "Success created",
			Data: models.CartResponse{
				ID:         int(cart.ID),
				Name:       cart.Shoes.Name,
				ShoesId:    int(cart.ShoesId),
				Price:      cart.Shoes.Price,
				Size:       cart.Size,
				Qty:        cart.Qty,
				TotalPrice: float64(cart.Qty) * cart.Shoes.Price,
				Status:     cart.Status,
				Images:     cart.Shoes.Images,
			},
		},
	)

}

func (cc *cartController) GetAllCarts(c echo.Context) error {
	userId := cc.authUsecase.ExtractTokenUserId(c, models.User_Type)
	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	carts, err := cc.cartUsecase.GetAllCarts(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	cartsResponse := make([]models.CartResponse, 0)

	for _, v := range carts {
		cartsResponse = append(cartsResponse, models.CartResponse{
			ID:         int(v.ID),
			Name:       v.Shoes.Name,
			ShoesId:    int(v.ShoesId),
			Price:      v.Shoes.Price,
			Size:       v.Size,
			Qty:        v.Qty,
			TotalPrice: float64(v.Qty) * v.Shoes.Price,
			Status:     v.Status,
			Images:     v.Shoes.Images,
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Message: "Success get all carts",
		Status:  http.StatusOK,
		Data:    cartsResponse,
	})
}

func (cc *cartController) UpdateCart(c echo.Context) error {
	var payload models.Carts

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userId := cc.authUsecase.ExtractTokenUserId(c, models.User_Type)
	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	cart, err := cc.cartUsecase.UpdateCart(id, userId, payload)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated,
		models.HttpResponse{
			Status:  http.StatusCreated,
			Message: "Success update cart",
			Data: models.CartResponse{
				ID:         int(cart.ID),
				Name:       cart.Shoes.Name,
				ShoesId:    int(cart.ShoesId),
				Price:      cart.Shoes.Price,
				Size:       cart.Size,
				Qty:        cart.Qty,
				TotalPrice: float64(cart.Qty) * cart.Shoes.Price,
				Status:     cart.Status,
				Images:     cart.Shoes.Images,
			},
		},
	)

}

func (cc *cartController) DeleteCartItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userId := cc.authUsecase.ExtractTokenUserId(c, models.User_Type)
	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	if err := cc.cartUsecase.DeleteCartItem(id); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.HttpResponse{
		Status:  http.StatusCreated,
		Message: "Success deleted",
	})
}
