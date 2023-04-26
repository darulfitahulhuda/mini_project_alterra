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
	CreateTransaction(payload dto.Transaction) (models.Transaction, error)
	GetAllTransaction() ([]models.Transaction, error)
	GetTransactionByUser(userId int) ([]models.Transaction, error)
	UpdateTransaction(id int, payload dto.Transaction) error
	UpdateStatusPayment(codePayment string, data string) error
	SoftDeleteTransaction(id int) error
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository) *transactionUsecase {
	return &transactionUsecase{transactionRepo: transactionRepo}
}

func (u *transactionUsecase) CreateTransaction(payload dto.Transaction) (models.Transaction, error) {
	var totalPrice float64
	var transactionDetails []models.TransactionDetail

	for _, v := range payload.Products {
		totalPrice += v.Price
		transactionDetails = append(transactionDetails, models.TransactionDetail{
			ShoesId: uint(v.ShoesId),
			Price:   v.Price,
			Qty:     v.Qty,
		})
	}

	transaction := models.Transaction{
		UserId:     uint(payload.UserId),
		TotalPrice: totalPrice,
		Status:     models.TRANSACTION_WAITING_PAYMENT,
		PaymentMethod: models.PaymentMethod{
			Name:        payload.NamePayment,
			CodePayment: randomCodePayment(payload.Products[0].ShoesId, payload.UserId),
			Status:      models.PAYMENT_STATUS_WAITING,
		},
		TransactionDetail: transactionDetails,
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

func (u *transactionUsecase) UpdateTransaction(id int, payload dto.Transaction) error {
	var totalPrice float64
	var transactionDetails []models.TransactionDetail

	for _, v := range payload.Products {
		totalPrice += v.Price
		transactionDetails = append(transactionDetails, models.TransactionDetail{
			ShoesId: uint(v.ShoesId),
			Price:   v.Price,
			Qty:     v.Qty,
		})
	}

	shipping := models.Shipping{
		Address:      payload.Shipping.Address,
		Price:        payload.Shipping.Price,
		Method:       payload.Shipping.Method,
		DeliveriDate: payload.Shipping.DeliveriDate,
	}

	transaction := models.Transaction{
		UserId:     uint(payload.UserId),
		TotalPrice: totalPrice,
		Status:     models.TRANSACTION_WAITING_PAYMENT,
		PaymentMethod: models.PaymentMethod{
			Name:   payload.NamePayment,
			Status: models.PAYMENT_STATUS_WAITING,
		},
		TransactionDetail: transactionDetails,
		Shipping:          shipping,
	}
	if err := u.transactionRepo.UpdateTransaction(id, transaction); err != nil {
		return err
	}

	return nil
}

func (u *transactionUsecase) UpdateStatusPayment(codePayment string, data string) error {
	if err := u.transactionRepo.UpdateStatusPayment(codePayment, data); err != nil {
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

	// Create a byte slice with length 5 to hold the random string
	randomCaracter := make([]byte, 5)

	for i := range randomCaracter {
		// Choose a random character from the character set
		randomCaracter[i] = charset[rand.Intn(len(charset))]
	}

	return strconv.Itoa(shoesId) + strconv.Itoa(userId) + string(randomCaracter)
}
