package repository

import (
	"main/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(data models.Transaction) (models.Transaction, error)
	GetAllTransaction() ([]models.Transaction, error)
	GetTransactionByUser(userId int) ([]models.Transaction, error)
	UpdateTransaction(id int, data models.Transaction) error
	UpdateStatusPayment(codePayment string, data string) error
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

func (r *transactionRepository) UpdateTransaction(id int, data models.Transaction) error {
	if err := r.db.Model(&models.Transaction{}).Where("ID = ?", id).Updates(data).Error; err != nil {
		return err
	}

	if err := r.db.Model(&models.TransactionDetail{}).Where("trasaction_id = ?", id).Updates(data.TransactionDetail).Error; err != nil {
		return err
	}

	if err := r.db.Model(&models.PaymentMethod{}).Where("trasaction_id = ?", id).Updates(data.PaymentMethod).Error; err != nil {
		return err
	}

	if err := r.db.Model(&models.Shipping{}).Where("trasaction_id = ?", id).Updates(data.Shipping).Error; err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) UpdateStatusPayment(codePayment string, data string) error {
	if err := r.db.Model(&models.PaymentMethod{}).Where("code_payment = ?", codePayment).Update("status", data).Error; err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) SoftDeleteTransaction(id int) error {
	if err := r.db.Where("ID = ?", id).Delete(&models.Transaction{}).Error; err != nil {
		return err
	}

	return nil
}
