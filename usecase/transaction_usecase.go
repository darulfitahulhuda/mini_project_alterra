package usecase

import (
	"main/dto"
	"main/models"
	"main/repository"
	"math/rand"
	"strconv"
	"time"
)

type TransactionUsecase interface {
	CreateTransaction(payload dto.TransactionRequest) (models.Transaction, error)
	GetAllTransaction() ([]models.Transaction, error)
	GetTransactionByUser(userId int) ([]models.Transaction, error)
	UpdateTransaction(id int, payload dto.TransactionRequest) (models.Transaction, error)
	UpdatePaymentMethod(payload dto.PaymentStatus) error
	SoftDeleteTransaction(id int) error
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
	shoesRepo       repository.ShoesRepository
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository, shoesRepo repository.ShoesRepository) *transactionUsecase {
	return &transactionUsecase{transactionRepo: transactionRepo, shoesRepo: shoesRepo}
}

func (u *transactionUsecase) CreateTransaction(payload dto.TransactionRequest) (models.Transaction, error) {
	var totalPrice float64
	var transactionDetails []models.TransactionDetail

	for _, v := range payload.Products {
		prodcutPrice := v.Price * float64(v.Qty)
		totalPrice += prodcutPrice
		transactionDetails = append(transactionDetails, models.TransactionDetail{
			ShoesId: uint(v.ShoesId),
			Price:   v.Price,
			Qty:     v.Qty,
			Size:    v.Size,
		})

		if err := u.shoesRepo.ReduceShoesQty(models.ShoesSize{ShoesId: uint(v.ShoesId), Size: v.Size, Qty: v.Qty}); err != nil {
			return models.Transaction{}, err

		}
	}

	totalPrice += payload.Shipping.Price

	shipping := models.Shipping{
		Address:      payload.Shipping.Address,
		Price:        payload.Shipping.Price,
		Method:       payload.Shipping.Method,
		DeliveriDate: payload.Shipping.DeliveriDate,
	}

	transaction := models.Transaction{
		UserId:     uint(payload.UserId),
		TotalPrice: totalPrice,
		Status:     models.TRANSACTION_PAYMENT_WAITING,
		PaymentMethod: models.PaymentMethod{
			Name:        payload.NamePayment,
			CodePayment: randomCodePayment(payload.Products[0].ShoesId, payload.UserId),
			Status:      models.PAYMENT_STATUS_WAITING,
		},
		TransactionDetail: transactionDetails,
		Shipping:          shipping,
	}

	data, err := u.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return data, err
	}

	return data, nil

}

func (u *transactionUsecase) GetAllTransaction() ([]models.Transaction, error) {
	transactions, err := u.transactionRepo.GetAllTransaction()
	if err != nil {
		return []models.Transaction{}, err

	}
	return transactions, nil
}

func (u *transactionUsecase) GetTransactionByUser(userId int) ([]models.Transaction, error) {
	transactions, err := u.transactionRepo.GetTransactionByUser(userId)

	if err != nil {
		return []models.Transaction{}, err

	}

	return transactions, nil

}

func (u *transactionUsecase) UpdateTransaction(id int, payload dto.TransactionRequest) (models.Transaction, error) {
	var totalPrice float64
	var transactionDetails []models.TransactionDetail

	for _, v := range payload.Products {
		transactionDetails = append(transactionDetails, models.TransactionDetail{
			ShoesId: uint(v.ShoesId),
			Price:   v.Price,
			Qty:     v.Qty,
			ID:      uint(v.ID),
		})
	}

	if err := u.transactionRepo.UpdateTransactionDetail(transactionDetails); err != nil {
		return models.Transaction{}, err
	}

	shipping := models.Shipping{
		Address:      payload.Shipping.Address,
		Price:        payload.Shipping.Price,
		Method:       payload.Shipping.Method,
		DeliveriDate: payload.Shipping.DeliveriDate,
	}

	if err := u.transactionRepo.UpdateShipping(id, shipping); err != nil {
		return models.Transaction{}, err

	}

	getTransaction, err := u.transactionRepo.GetTransactionById(id)
	if err != nil {
		return models.Transaction{}, err

	}

	for _, v := range getTransaction.TransactionDetail {
		prodcutPrice := v.Price * float64(v.Qty)
		totalPrice += prodcutPrice
	}

	totalPrice += getTransaction.Shipping.Price

	transaction := models.Transaction{
		UserId:     uint(payload.UserId),
		TotalPrice: totalPrice,
		Status:     payload.Status,
	}

	if err := u.transactionRepo.UpdateTransaction(id, transaction); err != nil {
		return models.Transaction{}, err

	}

	getTransactionId, err := u.transactionRepo.GetTransactionById(id)
	if err != nil {
		return models.Transaction{}, err

	}

	return getTransactionId, nil
}

func (u *transactionUsecase) UpdatePaymentMethod(payload dto.PaymentStatus) error {
	data := models.PaymentMethod{
		CodePayment: payload.CodePayment,
		Status:      models.PAYMENT_STATUS_SUCCESS,
	}
	paymentMethod, err := u.transactionRepo.UpdatePaymentMethod(data)
	if err != nil {
		return err
	}

	if err := u.transactionRepo.UpdateTransaction(int(paymentMethod.TransactionId), models.Transaction{Status: models.TRANSACTION_ADMIN_CONFIRMATION}); err != nil {
		return err
	}

	return nil
}
func (u *transactionUsecase) SoftDeleteTransaction(id int) error {
	if err := u.transactionRepo.SoftDeleteTransaction(id); err != nil {
		return err

	}

	return nil
}

func randomCodePayment(shoesId, userId int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Create a byte slice with length 15 to hold the random string
	randomCaracter := make([]byte, 15)

	for i := range randomCaracter {
		// Choose a random character from the character set
		randomCaracter[i] = charset[rand.Intn(len(charset))]
	}

	return strconv.Itoa(shoesId) + strconv.Itoa(userId) + string(randomCaracter)
}
