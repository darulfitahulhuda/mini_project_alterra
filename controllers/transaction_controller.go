package controllers

import (
	"main/dto"
	"main/models"
	"main/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionController interface{}

type transactionController struct {
	transactionUsecase usecase.TransactionUsecase
	authUsecase        usecase.AuthUsecase
}

func NewTransactionController(transactionUsecase usecase.TransactionUsecase, authUsecase usecase.AuthUsecase) *transactionController {
	return &transactionController{transactionUsecase: transactionUsecase, authUsecase: authUsecase}
}

func (t *transactionController) CreateTransaction(c echo.Context) error {
	var data dto.TransactionRequest

	userId := t.authUsecase.ExtractTokenUserId(c, models.User_Type)

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

	data.UserId = userId

	transaction, err := t.transactionUsecase.CreateTransaction(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	transactionResponse := changeToTransactionResponse(transaction)

	return c.JSON(http.StatusCreated, models.HttpResponse{
		Status:  http.StatusCreated,
		Message: "Success Craeted",
		Data:    transactionResponse,
	})
}

func (t *transactionController) GetAllTransaction(c echo.Context) error {
	userId := t.authUsecase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	transactions, err := t.transactionUsecase.GetAllTransaction()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	transactionResponse := changeToArrayTransactionResponse(transactions)

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success Get all transaction",
		Data:    transactionResponse,
	})

}

func (t *transactionController) GetTransactionByUser(c echo.Context) error {
	userId := t.authUsecase.ExtractTokenUserId(c, models.User_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	transactions, err := t.transactionUsecase.GetTransactionByUser(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	transactionResponse := changeToArrayTransactionResponse(transactions)

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success Get Transaction",
		Data:    transactionResponse,
	})

}

func (t *transactionController) UpdateTransaction(c echo.Context) error {
	var data dto.TransactionRequest

	userId := t.authUsecase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	transaction, erro := t.transactionUsecase.UpdateTransaction(id, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: erro.Error(),
		})

	}

	transactionResponse := changeToTransactionResponse(transaction)

	return c.JSON(http.StatusOK,
		models.HttpResponse{
			Status:  http.StatusOK,
			Message: "Success Updated",
			Data:    transactionResponse,
		},
	)
}

func (t *transactionController) UpdatePaymentMethod(c echo.Context) error {
	var data dto.PaymentStatus

	userId := t.authUsecase.ExtractTokenUserId(c, "all")

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

	err := t.transactionUsecase.UpdatePaymentMethod(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success Updated",
	})
}
func (t *transactionController) SoftDeleteTransaction(c echo.Context) error {
	userId := t.authUsecase.ExtractTokenUserId(c, models.User_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.HttpResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err2 := t.transactionUsecase.SoftDeleteTransaction(id)

	if err2 != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success Deleted",
	})
}
func changeToTransactionResponse(transaction models.Transaction) models.TransactionResponse {
	produts := make([]models.TransactionProductResponse, 0)
	for _, v := range transaction.TransactionDetail {
		produts = append(produts, models.TransactionProductResponse{
			ShoesId: v.ShoesId,
			Size:    v.Size,
			Qty:     v.Qty,
			Price:   v.Price,
		})

	}

	return models.TransactionResponse{
		ID:         int(transaction.ID),
		UserId:     int(transaction.UserId),
		TotalPrice: transaction.TotalPrice,
		Status:     transaction.Status,
		Products:   produts,
		Payment: models.PaymentMethodResponse{
			Name:        transaction.PaymentMethod.Name,
			CodePayment: transaction.PaymentMethod.CodePayment,
			Status:      transaction.PaymentMethod.Status,
		},
		Shipping: models.ShippingResponse{
			Address:      transaction.Shipping.Address,
			Price:        transaction.Shipping.Price,
			Name:         transaction.Shipping.Method,
			DeliveriDate: transaction.Shipping.DeliveriDate,
		},
	}
}

func changeToArrayTransactionResponse(transactions []models.Transaction) []models.TransactionResponse {
	transactionResponse := make([]models.TransactionResponse, 0)

	for _, transaction := range transactions {
		produts := make([]models.TransactionProductResponse, 0)
		for _, v := range transaction.TransactionDetail {
			produts = append(produts, models.TransactionProductResponse{
				ShoesId: v.ShoesId,
				Size:    v.Size,
				Qty:     v.Qty,
				Price:   v.Price,
			})

		}
		transactionResponse = append(transactionResponse, models.TransactionResponse{
			ID:         int(transaction.ID),
			UserId:     int(transaction.UserId),
			TotalPrice: transaction.TotalPrice,
			Status:     transaction.Status,
			Products:   produts,
			Payment: models.PaymentMethodResponse{
				Name:        transaction.PaymentMethod.Name,
				CodePayment: transaction.PaymentMethod.CodePayment,
				Status:      transaction.PaymentMethod.Status,
			},
			Shipping: models.ShippingResponse{
				Address:      transaction.Shipping.Address,
				Price:        transaction.Shipping.Price,
				Name:         transaction.Shipping.Method,
				DeliveriDate: transaction.Shipping.DeliveriDate,
			},
		})

	}
	return transactionResponse
}
