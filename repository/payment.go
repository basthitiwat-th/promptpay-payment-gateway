package repository

import (
	"time"

	"promptpay-payment-gateway/constants"

	"gorm.io/gorm"
)

type PaymentHistoryEntity struct {
	PaymentID     int       `gorm:"primaryKey;column:payment_id"`
	ReferenceID   string    `gorm:"colum:reference_id"`
	CustomerAccNo string    `gorm:"column:customer_acc_no"`
	PromptPayID   string    `gorm:"column:prompt_pay_id"`
	PaymentType   string    `gorm:"column:payment_type"`
	MerchantID    int       `gorm:"column:merchant_id"`
	Amount        float64   `gorm:"column:amount"`
	Status        string    `gorm:"colum:status"`
	CreatedAt     time.Time `gorm:"created_at"`
	UpdatedAt     time.Time `gorm:"updated_at"`
}

func (PaymentHistoryEntity) TableName() string {
	return "payment_history"
}

type PaymentStore struct {
	db *gorm.DB
}

func NewTransactionsStore(db *gorm.DB) *PaymentStore {
	return &PaymentStore{db}
}

func (r *PaymentStore) InsertOne(transaction PaymentHistoryEntity) error {
	if err := r.db.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func (r *PaymentStore) FindOne(reference_id string) (*PaymentHistoryEntity, error) {
	var paymentHistory PaymentHistoryEntity
	if err := r.db.Where("reference_id = ? and status = ?", reference_id, constants.STATUS_PENDING).First(&paymentHistory).Error; err != nil {
		return nil, err
	}

	return &paymentHistory, nil
}

func (r *PaymentStore) UpdateStatus(reference_id string) error {
	if err := r.db.Model(&PaymentHistoryEntity{}).Where("reference_id = ?", reference_id).Update("status", constants.STATUS_SUCCESS).Error; err != nil {
		return err
	}
	return nil
}
