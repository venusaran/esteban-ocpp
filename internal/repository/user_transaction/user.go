package usertransaction

import (
	"gorm.io/gorm"

	"github.com/Beep-Technologies/beepbeep3-ocpp/internal/models"
)

type BaseRepo interface {
	CreateUserTransaction(ut *models.UserTransaction) error
	UpdateUserTransaction(ut *models.UserTransaction) error
	GetUserTransactionByUserId(id int32) (txs []*models.UserTransaction, err error)
}

type baseRepo struct {
	db *gorm.DB
}

func NewBaseRepo(db *gorm.DB) BaseRepo {
	return &baseRepo{
		db: db,
	}
}

func (repo *baseRepo) CreateUserTransaction(ut *models.UserTransaction) error {
	return repo.db.Create(ut).Error
}

func (repo *baseRepo) UpdateUserTransaction(ut *models.UserTransaction) error {
	return repo.db.Save(ut).Error
}

func (repo *baseRepo) GetUserTransactionByUserId(id int32) (txs []*models.UserTransaction, err error) {
	var uts []*models.UserTransaction
	if err := repo.db.Where("user_id = ?", id).Find(&uts).Error; err != nil {
		return nil, err
	}
	return uts, nil
}
