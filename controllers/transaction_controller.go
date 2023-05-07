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
	var data dto.Transaction

	userId := t.authUsecase.ExtractTokenUserId(c, models.User_Type)

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

	data.UserId = userId

	transaction, err := t.transactionUsecase.CreateTransaction(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.TransactionResponse{
		Status:  http.StatusCreated,
		Message: "Success Craeted",
		Data:    transaction,
	})
}

func (t *transactionController) GetAllTransaction(c echo.Context) error {
	userId := t.authUsecase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	transactions, err := t.transactionUsecase.GetAllTransaction()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.TransactionListResponse{
		Status:  http.StatusOK,
		Message: "Success Get all transaction",
		Data:    transactions,
	})

}

func (t *transactionController) GetTransactionByUser(c echo.Context) error {
	userId := t.authUsecase.ExtractTokenUserId(c, models.User_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	transactions, err := t.transactionUsecase.GetTransactionByUser(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.TransactionListResponse{
		Status:  http.StatusOK,
		Message: "Success Get Transaction",
		Data:    transactions,
	})

}

func (t *transactionController) UpdateTransaction(c echo.Context) error {
	var data dto.Transaction

	userId := t.authUsecase.ExtractTokenUserId(c, models.Admin_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	transaction, erro := t.transactionUsecase.UpdateTransaction(id, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: erro.Error(),
		})

	}

	return c.JSON(http.StatusOK,
		models.HttpResponse{
			Status:  http.StatusOK,
			Message: "Success Updated",
			Data:    transaction,
		},
	)
}

func (t *transactionController) UpdatePaymentMethod(c echo.Context) error {
	var data dto.PaymentStatus

	userId := t.authUsecase.ExtractTokenUserId(c, "all")

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

	err := t.transactionUsecase.UpdatePaymentMethod(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success Updated",
	})
}
func (t *transactionController) SoftDeleteTransaction(c echo.Context) error {
	userId := t.authUsecase.ExtractTokenUserId(c, models.User_Type)

	if userId == 0 {
		return c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token Unauthorized",
			})
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err2 := t.transactionUsecase.SoftDeleteTransaction(id)

	if err2 != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success Deleted",
	})
}
