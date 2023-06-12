package service

import (
	"errors"
	"gorm.io/gorm"
	"placio-app/models"
)

type BusinessAccountService interface {
	CreateBusinessAccount(businessAccount *models.BusinessAccount) (*models.BusinessAccount, error)
	GetBusinessAccount(businessAccountID string) (*models.BusinessAccount, error)
	UpdateBusinessAccount(businessAccount *models.BusinessAccount) (*models.BusinessAccount, error)
	DeleteBusinessAccount(businessAccountID string) error
	ListBusinessAccounts(page, pageSize int, sortBy string, filters map[string]interface{}) ([]*models.BusinessAccount, error)
}

type BusinessAccountServiceImpl struct {
	store *models.BusinessAccount
	db    *gorm.DB
}

func NewBusinessAccountService(db *gorm.DB) BusinessAccountService {
	return &BusinessAccountServiceImpl{db: db, store: &models.BusinessAccount{}}
}

func (bas *BusinessAccountServiceImpl) CreateBusinessAccount(businessAccount *models.BusinessAccount) (*models.BusinessAccount, error) {
	if businessAccount == nil {
		return nil, errors.New("businessAccount cannot be nil")
	}

	if err := bas.db.Create(&businessAccount).Error; err != nil {
		return nil, err
	}

	return businessAccount, nil
}

func (bas *BusinessAccountServiceImpl) GetBusinessAccount(businessAccountID string) (*models.BusinessAccount, error) {
	if businessAccountID == "" {
		return nil, errors.New("businessAccountID cannot be nil")
	}

	var businessAccount models.BusinessAccount
	if err := bas.db.First(&businessAccount, "id = ?", businessAccountID).Error; err != nil {
		return nil, err
	}

	return &businessAccount, nil
}

func (bas *BusinessAccountServiceImpl) UpdateBusinessAccount(businessAccount *models.BusinessAccount) (*models.BusinessAccount, error) {
	if businessAccount == nil {
		return nil, errors.New("businessAccount cannot be nil")
	}

	if err := bas.db.Save(&businessAccount).Error; err != nil {
		return nil, err
	}

	return businessAccount, nil
}

func (bas *BusinessAccountServiceImpl) DeleteBusinessAccount(businessAccountID string) error {
	if businessAccountID == "" {
		return errors.New("businessAccountID cannot be nil")
	}

	if err := bas.db.Delete(&models.BusinessAccount{}, "id = ?", businessAccountID).Error; err != nil {
		return err
	}

	return nil
}

func (bas *BusinessAccountServiceImpl) ListBusinessAccounts(page, pageSize int, sortBy string, filters map[string]interface{}) ([]*models.BusinessAccount, error) {
	var businessAccounts []*models.BusinessAccount

	db := bas.db

	if len(filters) > 0 {
		db = db.Where(filters)
	}

	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Order(sortBy).Find(&businessAccounts).Error; err != nil {
		return nil, err
	}

	return businessAccounts, nil
}
