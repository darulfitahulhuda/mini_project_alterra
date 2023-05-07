package repository

import (
	"fmt"
	"main/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(data models.Transaction) (models.Transaction, error)

	GetAllTransaction() ([]models.Transaction, error)
	GetTransactionByUser(userId int) ([]models.Transaction, error)
	GetTransactionById(id int) (models.Transaction, error)

	UpdateTransaction(id int, data models.Transaction) error
	UpdateTransactionDetail(data []models.TransactionDetail) error
	UpdateShipping(id int, data models.Shipping) error
	UpdatePaymentMethod(data models.PaymentMethod) (models.PaymentMethod, error)

	SoftDeleteTransaction(id int) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(data models.Transaction) (models.Transaction, error) {

	if err := r.db.Create(&data).Error; err != nil {
		return models.Transaction{}, nil
	}

	return data, nil
}
func (r *transactionRepository) GetAllTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction

	if err := r.db.Preload("PaymentMethod").Preload("TransactionDetail").Preload("Shipping").Find(&transactions).Error; err != nil {
		return []models.Transaction{}, err
	}

	return transactions, nil
}

func (r *transactionRepository) GetTransactionByUser(userId int) ([]models.Transaction, error) {
	var transactions []models.Transaction

	if err := r.db.Where("user_id = ?", userId).Preload("PaymentMethod").Preload("TransactionDetail").Preload("Shipping").Find(&transactions).Error; err != nil {
		return []models.Transaction{}, err
	}

	return transactions, nil
}

func (r *transactionRepository) GetTransactionById(id int) (models.Transaction, error) {
	var transaction models.Transaction

	if err := r.db.Preload("PaymentMethod").Preload("TransactionDetail").Preload("Shipping").First(&transaction, id).Error; err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil

}

func (r *transactionRepository) UpdateTransaction(id int, data models.Transaction) error {
	if err := r.db.Model(&models.Transaction{}).Where("ID = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) UpdateTransactionDetail(data []models.TransactionDetail) error {

	for _, v := range data {
		var newData models.TransactionDetail

		if err := r.db.Model(&newData).Where("ID = ?", v.ID).Updates(v).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *transactionRepository) UpdateShipping(id int, data models.Shipping) error {
	var shipping models.Shipping
	if err := r.db.Model(&shipping).Where("trasaction_id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) UpdatePaymentMethod(data models.PaymentMethod) (models.PaymentMethod, error) {
	var payment models.PaymentMethod

	// if err := r.db.Model(&payment).Where("code_payment = ?", data.CodePayment).Updates(data).Error; err != nil {
	// 	return err
	// }
	if err := r.db.Where("code_payment = ?", data.CodePayment).First(&payment).Error; err != nil {
		return models.PaymentMethod{}, err
	}

	fmt.Printf("UpdatePaymentMethod: %d", payment.TransactionId)

	payment.Status = data.Status

	if err := r.db.Save(&payment).Error; err != nil {
		return models.PaymentMethod{}, err

	}

	return payment, nil
}

func (r *transactionRepository) SoftDeleteTransaction(id int) error {
	if err := r.db.Where("ID = ?", id).Delete(&models.Transaction{}).Error; err != nil {
		return err
	}

	return nil
}
