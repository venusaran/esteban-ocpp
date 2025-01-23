package user

import (
	"gorm.io/gorm"

	"github.com/Beep-Technologies/beepbeep3-ocpp/internal/models"
)

type BaseRepo interface {
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByHashedApiKey(hashedApiKey string) (*models.User, error)
}

type baseRepo struct {
	db *gorm.DB
}

func NewBaseRepo(db *gorm.DB) BaseRepo {
	return &baseRepo{
		db: db,
	}
}

func (repo *baseRepo) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *baseRepo) UpdateUser(user *models.User) error {
	return repo.db.Save(user).Error
}

func (repo *baseRepo) GetUserByUsername(username string) (*models.User, error) {
	var user *models.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *baseRepo) GetUserByHashedApiKey(hashedApiKey string) (*models.User, error) {
	var user *models.User
	if err := repo.db.Where("hashed_api_key = ?", hashedApiKey).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
